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
	StartSession
)

// Event defines any incoming event send by a client
// where the Type refers to the type of action which triggered the event
type Event interface {
	Json() ([]byte, error)
}

// RawClickEvent holds meta data about triggered click events
type RawClickEvent struct {
	AppUuid     string
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
	AppUuid     string
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
	AppUuid     string
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
	AppUuid        string    `cql:"app"`
	DeviceIP       string    `cql:"ip"`
	Type           EventType `json:"type"`
	Action         string    `json:"action" cql:"action"`
	FromStageLabel string    `json:"from_stage_label" cql:"from_stage_label"`
	FromStageID    string    `json:"from_stage" cql:"from_stage"`
	ToStageLabel   string    `json:"to_stage_label" cql:"to_stage_label"`
	ToStageID      string    `json:"to_stage" cql:"to_stage"`
	Timestamp      int64     `json:"timestamp" cql:"timestamp"`
	ElapsedTime    int64     `json:"elapsed_time" cql:"elapsed"`
}

func (evt FunnelChangeEvent) Json() ([]byte, error) {
	return json.Marshal(evt)
}

type SessionStart struct {
	AppUuid  string
	DeviceIP string
	Session  Record
}

func (evt SessionStart) Json() ([]byte, error) {
	return json.Marshal(evt)
}
