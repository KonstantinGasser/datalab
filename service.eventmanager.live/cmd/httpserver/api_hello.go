package httpserver

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.eventmanager.live/jwts"
	"github.com/sirupsen/logrus"
)

// Hello takes care about verifing that the user has the right to be tracked (lol sounds funny)
// It therfore performs all authentication and authrorization checks, loading meta data the client needs
func (server *Server) Hello(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[server.Hello] received request\n")

	wsTicket, errTicket := jwts.WebSocketTicket(r.Context().Value(typeKeyIP(keyIP)))
	if errTicket != nil {
		logrus.Errorf("<%v>[service.StartSession] could not issue ws ticket\n", r.Host)
		server.onErr(w, http.StatusInternalServerError, "could not issue ws ticket")
		return
	}
	server.onSuccess(w, http.StatusOK, map[string]interface{}{
		"ticket": wsTicket,
	})
}
