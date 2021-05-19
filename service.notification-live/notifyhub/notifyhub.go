package notifyhub

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type MessageEvent int

// VueMutation represents the actual function names
// on the vuejs client side. The "mutation" in a message will
// trigger the function on the client side
type VueMutation string

const (
	healthTimeOut = 20 * time.Second

	EventAppInvite MessageEvent = iota

	// MutationAppInvite maps to the corresponding Vue function
	// when the client socket receives a new message
	MutationAppInvite VueMutation = "APP_INVITE"
)

type Notification struct {
	Organization string
	Uuid         string
	Msg          Message
}

type Message struct {
	Mutation string                 `json:"mutation"`
	Event    MessageEvent           `json:"event"`
	Value    map[string]interface{} `json:"value"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type NotifyHub struct {
	subscribe   chan *Connection
	unsubscribe chan *Connection
	Notify      chan *Notification

	// guards the Organizations map
	mu sync.RWMutex
	// key is the organization name the client registered with
	Organizations map[string]*OrganizationPool
}

// New creates a new NotifyHub and starts its run function in a goroutine
func New() *NotifyHub {
	hub := &NotifyHub{
		subscribe:     make(chan *Connection),
		unsubscribe:   make(chan *Connection),
		Notify:        make(chan *Notification),
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
			hub.subscribeConn(conn)
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
			err := hub.sendMessage(notification)
			if err != nil {
				// receiver not available: kill connection
				if err == ErrWriteToConn {
					pool := hub.find(notification.Organization)
					if pool == nil {
						continue
					}
					hub.unsubscribe <- pool.Find(notification.Uuid)
				}
				logrus.Errorf("[notifyHub.chan.Notify] could not send message: %v\n", err)
			}
		// is this just for health checks
		case <-ticker.C:
			logrus.Infof("[notifyHub.run] status: running: (goroutines: %d) %+v\n", runtime.NumGoroutine(), hub.Organizations["datalab.dev"])
		}
	}
}

// sendMessage handles the sending of a message routing it to the correct pool
func (hub *NotifyHub) sendMessage(notification *Notification) error {
	hub.mu.RLock()
	defer hub.mu.RUnlock()

	pool, ok := hub.Organizations[notification.Organization]
	if !ok {
		return fmt.Errorf("could not find pool: %v", notification.Organization)
	}
	err := pool.Send(notification.Uuid, notification.Msg)
	if err != nil {
		return err
	}
	return nil
}

// subscribeConn adds a user the its correct pool
// if no pool present it creates a new one and adds the new
// connection to the pool
func (hub *NotifyHub) subscribeConn(conn *Connection) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	orgnPool, ok := hub.Organizations[conn.Organization]
	if !ok {
		hub.Organizations[conn.Organization] = &OrganizationPool{conn}
		return
	}
	if err := orgnPool.Add(conn); err != nil {
		logrus.Errorf("[notifyHub.subscribeConn] could not add connection: %v\n", err)
	}
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
