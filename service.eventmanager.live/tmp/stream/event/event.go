package event

type EventType int

const (
	RawClick EventType = iota
	RawURL
	StageURL
	StageClick
)

type Event struct {
	// Type indicated what origin the event has. see const for Type
	Type EventType
}

// RawClickEvent represents any click a user performs
type RawClickEvent struct {
	Event
	Timestamp  int64
	CurrentURL string
	HTMLTarget string
	// ElapsedTime stands for the time since the last RawClickEvent
	ElapsedTime int64
}

// RawURLEvent represents any URL change. The event stores the url the moved from
// the url the user went to and the elapsed time the user spend on the "From" URL
type RawURLEvent struct {
	Event
	Timestamp   int64
	From        string
	To          string
	ElapsedTime int64
}
