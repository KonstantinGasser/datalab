package bus

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/KonstantinGasser/datalab/service.eventmanager.live/internal/session"
	"github.com/KonstantinGasser/datalab/service.eventmanager.live/ports/cassandra"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// matchOriginFromToken checks if the request origin matches with the app-token origin
// alloing to block connection from unwanted clients
var matchOriginFromToken = func(r *http.Request) bool {
	return true
	// allowedOrigin := r.Header.Get("Origin")
	// requestedOrigin := ctx_value.GetString(r.Context(), "app.origin")

	// return len(allowedOrigin) != 0 && requestedOrigin == allowedOrigin
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     matchOriginFromToken,
}

type PubSub struct {
	// sub creates a new active session based on the session.Record
	// sub chan *session.User
	// publish publishlishes an incoming event, distributing it to the right sink
	publish chan session.Event
	cqlC    *cassandra.Client
	// drop deletes lost or droped connections, finishing a session
	drop chan *session.User // ip:port
	// // mu garudes session map
	// mu      sync.RWMutex
	// session map[string]*session.User
}

// NewPubSub returns a new instance of a PubSub
func NewPubSub(cqlC *cassandra.Client) *PubSub {
	return &PubSub{
		// sub:  make(chan *session.User),
		publish: make(chan session.Event),
		drop:    make(chan *session.User),
		cqlC:    cqlC,

		// session: make(map[string]*session.User),
	}
}

// Start starts N hub to listen to incoming events and handle protocoll upgrades
func (hub *PubSub) Start(scaler int) {
	for i := 1; i <= scaler; i++ {
		go hub.run(i)
	}
}

// run runs the event-loop in its own goroutine handling incoming events
func (hub *PubSub) run(workerID int) {
	// ticker for health logs
	var ticker = time.NewTicker(30 * time.Second)

	for {
		select {
		case user := <-hub.drop:
			err := hub.cqlC.InsertEvent(
				context.Background(),
				cassandra.InsertSession,
				user.AppUuid,
				time.Now(),
				user.DeviceIP,
				user.Start,
				user.Duration,
				user.Record.Referrer,
			)
			if err != nil {
				logrus.Errorf("[event-bus.run] could not persist InsertSession: %v\n", err)
			}
		case event := <-hub.publish:
			fmt.Printf("EVENT: %T, %+v\n", event, event)
			switch evt := event.(type) {
			case session.SessionStart:
				err := hub.cqlC.InsertEvent(
					context.Background(),
					cassandra.InsertBatchStart,
					evt.AppUuid,
					evt.Session.Device,
					evt.AppUuid,
					evt.Session.Browser,
					evt.AppUuid,
					time.Now(),
				)
				if err != nil {
					logrus.Errorf("[event-bus.run] could not persist InsertBatchStart: %v\n", err)
				}
			case session.RawURLEvent:
				err := hub.cqlC.InsertEvent(context.Background(),
					cassandra.UpsertPageHit,
					evt.ElapsedTime,
					evt.AppUuid,
					evt.From,
					evt.Timestamp,
				)
				if err != nil {
					logrus.Errorf("[event-bus.run] could not persist UpsertPageHit: %v\n", err)
				}
			case session.BtnTimeEvent:
				var deltaElapsedClick, deltaElapsedLeave, deltaHoverClick, deltaHoverLeave int
				deltaElapsedClick = int(evt.ElapsedTime)
				deltaHoverClick = 1
				if evt.Action == "hover-leave" {
					deltaElapsedClick = 0
					deltaHoverClick = 0

					deltaElapsedLeave = int(evt.ElapsedTime)
					deltaHoverLeave = 1
				}
				err := hub.cqlC.InsertEvent(context.Background(),
					cassandra.UpsertInterestingBtn,
					deltaElapsedClick,
					deltaElapsedLeave,
					deltaHoverClick,
					deltaHoverLeave,
					evt.AppUuid,
					evt.Target,
				)
				if err != nil {
					logrus.Errorf("[event-bus.run] could not persist UpsertInterestingBtn: %v\n", err)
				}
			case session.FunnelChangeEvent:
				if err := hub.cqlC.InsertEvent(
					context.Background(),
					cassandra.InsertFunnelChange,
					evt.AppUuid,
					evt.Timestamp,
					evt.ToStageLabel,
					evt.DeviceIP,
				); err != nil {
					logrus.Errorf("[event-bus.run] could not persist InsertFunnelChange: %v\n", err)
				}
				if err := hub.cqlC.InsertEvent(context.Background(),
					cassandra.UpsertAvgTimeStage,
					evt.ElapsedTime,
					evt.AppUuid,
					evt.FromStageLabel,
				); err != nil {
					logrus.Errorf("[event-bus.run] could not persist UpsertAvgTimeStage: %v\n", err)
				}
			}
		case <-ticker.C:
			logrus.Infof("[event-bus.run-%d] goroutines: %d\n", workerID, runtime.NumGoroutine())
		}
	}
}

// UpgradeProtocoll hajiecks the ResponseWriter and upgrdares the http protocoll to a ws
// if successful starts the listener on the connection.
func (hub *PubSub) UpgradeProtocoll(w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	user, err := session.NewUser(r, conn, hub.publish, hub.drop)
	if err != nil {
		return err
	}
	hub.publish <- session.SessionStart{
		AppUuid:  user.AppUuid,
		DeviceIP: user.DeviceIP,
		Session:  user.Record,
	}
	go user.Listen()
	return nil
}
