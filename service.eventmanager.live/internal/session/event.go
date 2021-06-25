package session

import (
	"encoding/json"
	"fmt"
)

var (
	ErrNoSuchType = fmt.Errorf("unknown event type")
)

type EventType int

const (
	// RawClick used for any click-event
	RawClick EventType = iota
	// RawURL used for any URL change
	RawURL
)

// Event defines any incoming event send by a client
// where the Type refers to the type of action which triggered the event
type Event interface {
	Json() ([]byte, error)
}

// RawClickEvent holds meta data about triggered click events
type RawClickEvent struct {
	Type        EventType
	CurrentURL  string
	Action      string
	Target      string
	Timestamp   int64
	ElapsedTime int64
}

func (evt RawClickEvent) Json() ([]byte, error) {
	return json.Marshal(evt)
}

// RawURLEvent holds meta data about triggered URL changes
type RawURLEvent struct {
	Type        EventType
	From        string
	To          string
	Action      string
	Timestamp   int64
	ElapsedTime int64
}

func (evt RawURLEvent) Json() ([]byte, error) {
	return json.Marshal(evt)
}
