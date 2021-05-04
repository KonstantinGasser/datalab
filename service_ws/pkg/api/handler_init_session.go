package api

import (
	"encoding/json"
	"net/http"
)

func (api API) HandlerInitSession(w http.ResponseWriter, r *http.Request) {
	var session struct {
		Meta struct {
			OS struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"OS"`
			Browser string `json:"browser"`
			Device  string `json:"device"`
		} `json:"meta"`
		Referrer string `json:"referrer"`
	}
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		api.onErr(w, http.StatusBadRequest, "could not decode r.Body")
		return
	}
	// fmt.Printf("Session-Record:\n%+v\n", session)
	w.WriteHeader(http.StatusOK)
}
