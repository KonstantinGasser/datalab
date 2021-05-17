package types

// SessionStart represents the initial data
// a client send before opening the web-socket
type SessionStart struct {
	AppUuid   string
	AppOrigin string
	Cookie    string
	Meta      struct {
		OS struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"OS"`
		Browser string `json:"browser"`
		Device  string `json:"device"`
	} `json:"meta"`
	Referrer string `json:"referrer"`
}
