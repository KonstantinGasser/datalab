package sessions

type Session struct {
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

type StartRequest struct {
	AppUuid, AppOrigin string
	Session            *Session
}
