package bus

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.eventmanager.live/internal/session"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// originFromAppToken checks if the request origin matches with the app-token origin
// alloing to block connection from unwanted clients
var originFromAppToken = func(r *http.Request) bool {
	allowedOrigin := r.Header.Get("Origin")
	requestedOrigin := ctx_value.GetString(r.Context(), "app.origin")

	return len(allowedOrigin) != 0 && requestedOrigin == allowedOrigin
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     originFromAppToken,
}

type PubSub struct {
	// sub creates a new active session based on the session.Record
	sub chan *session.User
	// pub publishes an incoming event, distributing it to the right sink
	pub chan session.Event
	// drop deletes lost or droped connections, finishing a session
	drop chan session.User // ip:port
	// mu garudes session map
	mu      sync.RWMutex
	session map[string]*session.User
}

// NewPubSub returns a new instance of a PubSub
func NewPubSub() *PubSub {
	return &PubSub{
		sub:  make(chan *session.User),
		pub:  make(chan session.Event),
		drop: make(chan session.User),

		session: make(map[string]*session.User),
	}
}

// Start starts all goroutines and effectivly allowing to connect and interact with the
// PupSub
func (hub *PubSub) Start() {
	hub.run()
}

// run runs the event-loop in its own goroutine handling incoming events
func (hub *PubSub) run() {
	// ticker for health logs
	var ticker = time.NewTicker(15 * time.Second)

	for {
		select {
		case user := <-hub.sub:
			fmt.Printf("RECORD: %v\n", *user)
			hub.mu.Lock()
			if _, ok := hub.session[user.IpPort]; ok {
				hub.mu.Unlock()
				continue
			}
			hub.session[user.IpPort] = user
			hub.mu.Unlock()
		case user := <-hub.drop:
			logrus.Infof("[event-bus.drop] deleting for %s\n", user.IpPort)
			hub.mu.Lock()
			delete(hub.session, user.IpPort)
			hub.mu.Unlock()
		case event := <-hub.pub:
			fmt.Printf("EVENT: %v\n", event)
		case <-ticker.C:
			logrus.Infof("[event-bus.run] connections: %d - goroutines: %d\n", len(hub.session), runtime.NumGoroutine())
		}
	}
}

func (hub *PubSub) UpgradeProtocoll(w http.ResponseWriter, r *http.Request) error {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	user, err := session.NewUser(r, conn, hub.pub, hub.drop)
	if err != nil {
		return err
	}
	go user.Listen()
	hub.sub <- user

	return nil
}
