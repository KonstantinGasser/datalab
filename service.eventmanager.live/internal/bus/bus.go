package bus

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/KonstantinGasser/datalab/service.eventmanager.live/internal/session"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
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
	go hub.run()
}

// run runs the event-loop in its own goroutine handling incoming events
func (hub *PubSub) run() {

	for {
		select {
		case user := <-hub.sub:
			fmt.Printf("RECORD: %v\n", *user)
		case event := <-hub.pub:
			fmt.Printf("EVENT: %v\n", event)
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
