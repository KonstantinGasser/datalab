package stream

import (
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.eventmanager.live/internal/stream/connection"
	"github.com/KonstantinGasser/datalab/service.eventmanager.live/ports/cassandra"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Sink string

const (
	Cassandra Sink = "cassandra"
)

type Stream struct {
	cassandra *cassandra.Client
	// subscribe, unsubscribe handles new clients
	// where each message represents a client uuid
	subscribe   chan string
	unsubscribe chan string

	PushIncident chan map[string]interface{}
}

func New(cc *cassandra.Client) *Stream {
	return &Stream{
		cassandra:    cc,
		subscribe:    make(chan string),
		unsubscribe:  make(chan string),
		PushIncident: make(chan map[string]interface{}),
	}
}

func (s *Stream) Listen() {
	logrus.Info("[stream.Listen] starting to listen for events\n")
	for {
		select {
		case clientUuid := <-s.subscribe:
			fmt.Println("Sub - CLIENT: ", clientUuid)
		case clientUuid := <-s.unsubscribe:
			fmt.Println("UnSub - CLIENT", clientUuid)
		case event := <-s.PushIncident:
			// - eval event by type
			// - assert to correct type
			// - send event to downstream sink
			fmt.Println("EVENT: ", event)
		}
	}
}

const keyCookie string = "x-datalab-cookie"

func (stream *Stream) HttpUpgrade(w http.ResponseWriter, r *http.Request) error {
	wsconn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	cookie, err := r.Cookie(keyCookie)
	if err != nil {
		return err
	}

	connection := connection.New(wsconn, stream.PushIncident, stream.unsubscribe, cookie.Name)
	go connection.Listen()

	return nil
}
