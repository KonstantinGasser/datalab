package notifyhub

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/KonstantinGasser/datalab/service.notification-live/repo"
	"github.com/sirupsen/logrus"
)

const (
	healthTimeOut = 10 * time.Second
)

type NotifyHub struct {
	repo         repo.Repo
	subscribe    chan *Connection
	unsubscribe  chan *Connection
	Notify       chan *IncomingEvent
	RemoveNotify chan *RemoveEvent
	HideNotify   chan *HideEvent
	batchNotify  chan *UserNotifications

	// guards the Organizations map
	mu sync.RWMutex
	// key is the organization name the client registered with
	Organizations map[string]*OrganizationPool
}

// New creates a new NotifyHub and starts its run function in a goroutine
func New(repo repo.Repo) *NotifyHub {
	hub := &NotifyHub{
		repo:          repo,
		subscribe:     make(chan *Connection),
		unsubscribe:   make(chan *Connection),
		Notify:        make(chan *IncomingEvent),
		RemoveNotify:  make(chan *RemoveEvent),
		batchNotify:   make(chan *UserNotifications),
		HideNotify:    make(chan *HideEvent),
		mu:            sync.RWMutex{},
		Organizations: make(map[string]*OrganizationPool),
	}
	go hub.run()
	return hub
}

// run runs in its own goroutine receiving register requests and handling
// notifications
func (hub *NotifyHub) run() {
	logrus.Infof("[notifyHub.run] status: starting\n")
	var ticker = time.NewTicker(healthTimeOut)
	for {
		select {
		// register adds a user the its correct pool
		// if no pool present it creates a new one and adds the new
		// connection to the pool
		case conn := <-hub.subscribe:
			err := hub.subscribeConn(conn)
			if err != nil {
				logrus.Errorf("[notifyHub.chan.subscribe] could not subscribe connection: %v\n", err)
			}
			// check health of connection
			// if unhealthy kill connection
			go conn.Health(hub)
			logrus.Infof("[notifyHub.run] registered: <%s> in <%s>\n", conn.Uuid, conn.Organization)
		// will remove connection form pool and if pool empty
		// delete the pool
		case conn := <-hub.unsubscribe:
			hub.unsubscribeConn(conn)
		// finds the correct pool to send the message
		// if send fails looks up connection and kills it
		case notification := <-hub.Notify:
			switch notification.Event {
			case EventAppInvite, EventAppInviteReminder:
				err := hub.sendMessage(notification)
				if err != nil {
					// receiver not available: kill connection
					if err == ErrWriteToConn {
						pool := hub.find(notification.Organization)
						if pool == nil {
							continue
						}
						hub.unsubscribe <- pool.Find(notification.UserUuid)
					}
					logrus.Errorf("[notifyHub.chan.Notify] could not send invite: %v\n", err)
				}
			case EventSyncApp:
				err := hub.sendBroadcast(notification)
				if err != nil {
					fmt.Println(notification)
					logrus.Errorf("[notifyHub.chan.Notify] could not broadcast message: %v\n", err)
				}
			}

		case userNotifies := <-hub.batchNotify:
			err := hub.sendBatch(userNotifies.Organization, userNotifies.UserUuid, userNotifies.Notifications)
			if err != nil {
				// receiver not available: kill connection
				if err == ErrWriteToConn {
					pool := hub.find(userNotifies.Organization)
					if pool == nil {
						continue
					}
					hub.unsubscribe <- pool.Find(userNotifies.UserUuid)
				}
				logrus.Errorf("[notifyHub.chan.batchNotify] could not send message: %v\n", err)
			}
			// removes notifications with are no longer important
			// and can be delete from the database
		case notification := <-hub.HideNotify:
			err := hub.HideNotification(notification)
			if err != nil {
				logrus.Errorf("[notifyHub.chan.HideNotify] could not hide message: %v\n", err)
			}
		// removes notifications with are no longer important
		// and can be delete from the database
		case notification := <-hub.RemoveNotify:
			err := hub.removeNotification(notification)
			if err != nil {
				logrus.Errorf("[notifyHub.chan.RemoveNotify] could not remove message: %v\n", err)
			}
		// is this just for health checks
		case <-ticker.C:
			logrus.Infof("[notifyHub.run] status: running: (goroutines: %d) %+v\n", runtime.NumGoroutine(), hub.Organizations["datalab.dev"])
		}
	}
}

// sendMessage handles the sending of a message routing it to the correct pool
func (hub *NotifyHub) sendMessage(notify *IncomingEvent) error {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	// save message in database
	go hub.SaveEvent(notify.UserUuid, notify)

	pool, ok := hub.Organizations[notify.Organization]
	if !ok {
		return fmt.Errorf("could not find pool: %v", notify.Organization)
	}
	err := pool.Send(notify.UserUuid, notify)
	if err != nil {
		return err
	}
	return nil
}

// sendBroadcast broadcasts a message to all connection in a given pool
func (hub *NotifyHub) sendBroadcast(notify *IncomingEvent) error {
	pool, ok := hub.Organizations[notify.Organization]
	if !ok {
		return fmt.Errorf("could not find pool: %v", notify.Organization)
	}
	err := pool.Broadcast(notify)
	if err != nil {
		return err
	}
	return nil
}

// sendMessage handles the sending of a message routing it to the correct pool
func (hub *NotifyHub) sendBatch(organization, userUuid string, messages []Notification) error {
	defer func() {
		// after delivery to client
		// save notification in storage
	}()
	hub.mu.Lock()
	defer hub.mu.Unlock()

	pool, ok := hub.Organizations[organization]
	if !ok {
		return fmt.Errorf("could not find pool: %v", organization)
	}
	err := pool.SendBatch(userUuid, BatchNotification{
		Mutation: MutationLoadFromDB,
		Messages: messages,
	})
	if err != nil {
		return err
	}
	return nil
}

// subscribeConn adds a user the its correct pool
// if no pool present it creates a new one and adds the new
// connection to the pool
func (hub *NotifyHub) subscribeConn(conn *Connection) error {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	ok, err := hub.HasRecord(conn.Uuid)
	if err != nil { // db error will lead to no socket connection
		return err
	}
	if !ok { // create default record for new client if non present
		err := hub.PersistInitRecord(conn)
		if err != nil {
			return err
		}
	}
	// add connection to correct pool
	orgnPool, ok := hub.Organizations[conn.Organization]
	if !ok {
		hub.Organizations[conn.Organization] = &OrganizationPool{conn}
	} else {
		if err := orgnPool.Add(conn); err != nil {
			return err
		}
	}

	// send persisted notification in background
	go hub.LookUpAndSend(conn.Uuid)
	return nil
}

func (hub *NotifyHub) unsubscribeConn(conn *Connection) {
	hub.mu.Lock()
	defer hub.mu.Unlock()

	if conn == nil {
		return
	}

	orgnPool, ok := hub.Organizations[conn.Organization]
	if !ok {
		logrus.Errorf("[notifyHub.unsubscribeConn] could not remove connection: no organization pool found\n")
		return
	}
	if err := orgnPool.Remove(conn); err != nil {
		logrus.Errorf("[notifyHub.unsubscribeConn] could not remove connection: %v\n", err)
		return
	}
	if orgnPool.Length() == 0 {
		delete(hub.Organizations, conn.Organization)
	}
}

func (hub *NotifyHub) removeNotification(notification *RemoveEvent) error {
	return hub.Remove(notification)
}

func (hub *NotifyHub) HideNotification(notification *HideEvent) error {
	return hub.Hide(notification)
}

func (hub *NotifyHub) find(orgn string) *OrganizationPool {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	return hub.Organizations[orgn]
}

func (hub *NotifyHub) Stop() {
	close(hub.subscribe)
	close(hub.unsubscribe)
	close(hub.Notify)
}

func (hub *NotifyHub) OpenSocket(ctx context.Context, w http.ResponseWriter, r *http.Request, uuid, orgn string) error {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("[hub.OpenSocket] could not open socket: %v\n", err)
		return err
	}
	connection := &Connection{
		Conn:         conn,
		Uuid:         uuid,
		Organization: orgn,
	}
	hub.subscribe <- connection
	return nil
}
