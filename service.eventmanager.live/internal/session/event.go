package session

import (
	"encoding/json"
	"fmt"
)

var (
	ErrNoSuchType  = fmt.Errorf("unknown event type")
	ErrNoTypeFound = fmt.Errorf("no event type found")
)

type EventType int

const (
	// RawClick used for any click-event
	RawClick EventType = iota
	// RawURL used for any URL change
	RawURL
	// BtnTime for configured btns
	BtnTime
	// FunnelChange for when the stage changes
	FunnelChange
)

// Event defines any incoming event send by a client
// where the Type refers to the type of action which triggered the event
type Event interface {
	Json() ([]byte, error)
}

// RawClickEvent holds meta data about triggered click events
type RawClickEvent struct {
	DeviceIP    string
	Type        EventType `json:"type"`
	CurrentURL  string    `json:"current_url"`
	Target      string    `json:"target"`
	Timestamp   int64     `json:"timestamp"`
	ElapsedTime int64     `json:"elapsed_time"`
}

func (evt RawClickEvent) Json() ([]byte, error) {
	return json.Marshal(evt)
}

// RawURLEvent holds meta data about triggered URL changes
type RawURLEvent struct {
	DeviceIP    string
	Type        EventType `json:"type"`
	From        string    `json:"from"`
	To          string    `json:"to"`
	Timestamp   int64     `json:"timestamp"`
	ElapsedTime int64     `json:"elapsed_time"`
}

func (evt RawURLEvent) Json() ([]byte, error) {
	return json.Marshal(evt)
}

// RawURLEvent holds meta data about triggered URL changes
type BtnTimeEvent struct {
	DeviceIP    string
	Type        EventType `json:"type"`
	Target      string    `json:"target"`
	Action      string    `json:"action"`
	Timestamp   int64     `json:"timestamp"`
	ElapsedTime int64     `json:"elapsed_time"`
}

func (evt BtnTimeEvent) Json() ([]byte, error) {
	return json.Marshal(evt)
}

// RawURLEvent holds meta data about triggered URL changes
type FunnelChangeEvent struct {
	DeviceIP    string
	Type        EventType `json:"type"`
	Action      string    `json:"action"`
	Entered     int       `json:"entered"`
	ElapsedTime int64     `json:"elapsed_time"`
}

func (evt FunnelChangeEvent) Json() ([]byte, error) {
	return json.Marshal(evt)
}
