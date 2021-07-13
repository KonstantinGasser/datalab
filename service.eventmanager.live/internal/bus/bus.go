package bus

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.eventmanager.live/internal/session"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// matchOriginFromToken checks if the request origin matches with the app-token origin
// alloing to block connection from unwanted clients
var matchOriginFromToken = func(r *http.Request) bool {
	allowedOrigin := r.Header.Get("Origin")
	requestedOrigin := ctx_value.GetString(r.Context(), "app.origin")

	return len(allowedOrigin) != 0 && requestedOrigin == allowedOrigin
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
	// drop deletes lost or droped connections, finishing a session
	// drop chan session.User // ip:port
	// // mu garudes session map
	// mu      sync.RWMutex
	// session map[string]*session.User
}

// NewPubSub returns a new instance of a PubSub
func NewPubSub() *PubSub {
	return &PubSub{
		// sub:  make(chan *session.User),
		publish: make(chan session.Event),
		// drop: make(chan session.User),

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
	var ticker = time.NewTicker(15 * time.Second)

	for {
		select {
		case event := <-hub.publish:
			fmt.Printf("EVENT: %+v\n", event)
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
	user, err := session.NewUser(r, conn, hub.publish)
	if err != nil {
		return err
	}
	go user.Listen()
	return nil
}
