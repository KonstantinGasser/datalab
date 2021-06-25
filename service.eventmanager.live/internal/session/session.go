package session

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

const (
	QueryRef       = "ref"
	QueryOsName    = "os_name"
	QueryOsVersion = "os_vs"
	QueryDevice    = "device"
	QueryBrowser   = "browser"
)

var (
	ErrNoIpPort = fmt.Errorf("ip:port pair is not valid")
)

// Record defines the initialization data to start tracking user sessions
type Record struct {
	OS struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"OS"`
	Browser  string `json:"browser"`
	Device   string `json:"device"`
	Referrer string `json:"referrer"`
}

type User struct {
	IpPort string
	// AppUuid refers to the app the user is located on
	AppUuid string
	// MappedOrgn refers to the organization in the system
	// to which the App is mapped to
	MappedOrgn string
	Record     Record

	connection *websocket.Conn

	publish chan<- Event
	drop    chan<- User
}

func NewUser(req *http.Request, conn *websocket.Conn, pub chan<- Event, drop chan<- User) (*User, error) {
	if !isIpPort(req.RemoteAddr) {
		return nil, ErrNoIpPort
	}
	user := User{
		IpPort:     req.RemoteAddr,
		connection: conn,
		publish:    pub,
		drop:       drop,
	}
	user.RecordFromURL(req.URL)
	return &user, nil
}

// Listen runs in its own goroutine listening for incoming user events
// and sending them to the publish chan
func (u User) Listen() {
	fmt.Println("Listen")
	defer func() {
		logrus.Infof("[user.Listen] closing for %s\n", u.IpPort)
		// wraping up user session sending before deletion
		u.drop <- u.finish()
	}()

	for {
		msgType, bytes, err := u.connection.ReadMessage()
		if err != nil || msgType == websocket.CloseMessage {
			break
		}
		event, err := u.Event(bytes)
		if err != nil {
			logrus.Errorf("[user.Listen] could not marshal bytes to proper event: %v\n", err)
			continue
		}
		u.publish <- event
	}
}

// Event converts the given bytes slice to a session known event type
func (u User) Event(bytes []byte) (Event, error) {
	var rawEvent map[string]interface{}
	if err := json.Unmarshal(bytes, &rawEvent); err != nil {
		return nil, err
	}

	evtType, ok := rawEvent["type"]
	if !ok {
		return nil, ErrNoTypeFound
	}
	_type, ok := evtType.(float64)
	if !ok {
		return nil, ErrNoSuchType
	}

	switch EventType(_type) {
	case RawClick:
		var evt RawClickEvent
		if err := json.Unmarshal(bytes, &evt); err != nil {
			return nil, err
		}
		return evt, nil
	case RawURL:
		var evt RawURLEvent
		if err := json.Unmarshal(bytes, &evt); err != nil {
			return nil, err
		}
		return evt, nil
	default:
		return nil, ErrNoSuchType
	}
}

// RecordFromURL queries for record data in the URL and assigns record
// to user
func (u *User) RecordFromURL(url *url.URL) {
	var r Record
	r.Referrer = url.Query().Get(QueryRef)
	r.OS.Name = url.Query().Get(QueryOsName)
	r.OS.Version = url.Query().Get(QueryOsVersion)
	r.Device = url.Query().Get(QueryDevice)
	r.Browser = url.Query().Get(QueryBrowser)

	u.Record = r
}

// finish: not yet sure what it needs to do but pretty sure I will need it
func (u User) finish() User {
	return u
}

// isIpPort checks if a given r.RemoteHost is a correct IP:PORT pair
func isIpPort(ipPort string) bool {
	parts := strings.Split(ipPort, ":")
	return len(parts) == 2
}
