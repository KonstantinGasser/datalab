package notifyhub

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type MessageEvent int

// VueMutation represents the actual function names
// on the vuejs client side. The "mutation" in a message will
// trigger the function on the client side
type VueMutation string

const (
	// indicates that a user has been invited to join an App
	EventAppInvite MessageEvent = iota
	// indicates that the App has changed and can be synced
	EventSyncApp

	// MutationAppInvite maps to the corresponding Vue function
	// when the client socket receives a new message
	MutationAppInvite  VueMutation = "APP_INVITE"
	MutationLoadFromDB VueMutation = "INIT_LOAD"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Notification represents one notification as it will
// be stored in the database.
// Each user will have its own information with a slice of messages
type UserNotifications struct {
	UserUuid      string         `bson:"_id"`
	Organization  string         `bson:"organization"`
	Notifications []Notification `bson:"notifications"`
}

// Message represents a message as it will be streamed
// to the client socket
type Notification struct {
	Timestamp int64                  `json:"timestamp"`
	Mutation  VueMutation            `json:"mutation" bson:"mutation"`
	Event     MessageEvent           `json:"event" bson:"event"`
	Value     map[string]interface{} `json:"value" bson:"value"`
}
type BatchNotification struct {
	Mutation VueMutation    `json:"mutation"`
	Messages []Notification `json:"notifications"`
}

type IncomingEvent struct {
	UserUuid     string                 `json:"receiver_uuid"`
	Organization string                 `json:"receiver_orgn"`
	Timestamp    int64                  `json:"timestamp"`
	Mutation     VueMutation            `json:"mutation"`
	Event        MessageEvent           `json:"event"`
	Value        map[string]interface{} `json:"value"`
}

type RemoveEvent struct {
	UserUuid  string `json:"user_uuid"`
	Timestamp int64  `json:"timestamp"`
}
