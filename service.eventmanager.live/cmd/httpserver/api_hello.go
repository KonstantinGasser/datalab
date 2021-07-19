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

	ctxOrigin := r.Context().Value(typeKeyClaims(keyOrigin))
	origin, ok := ctxOrigin.(string)
	if !ok {
		logrus.Errorf("<%v>[service.StartSession] origin of app-token not valid\n", r.Host)
		server.onErr(w, http.StatusUnauthorized, "app-token tamppered")
		return
	}
	ctxAppUuid := r.Context().Value(typeKeyClaims(keyAppUuid))
	uuid, ok := ctxAppUuid.(string)
	if !ok {
		logrus.Errorf("<%v>[service.StartSession] uuid of app-token not valid\n", r.Host)
		server.onErr(w, http.StatusUnauthorized, "app-token tamppered")
		return
	}
	if len(origin) == 0 {
		logrus.Errorf("<%v>[service.StartSession] origin of app-token not valid - length violation\n", r.Host)
		server.onErr(w, http.StatusUnauthorized, "app-token tamppered")
		return
	}

	wsTicket, errTicket := jwts.WebSocketTicket(r.Context().Value(typeKeyIP(keyIP)), origin, uuid)
	if errTicket != nil {
		logrus.Errorf("<%v>[service.StartSession] could not issue ws ticket\n", r.Host)
		server.onErr(w, http.StatusInternalServerError, "could not issue ws ticket")
		return
	}
	config, err := server.appConfigClient.GetAppConfig(r.Context(), uuid)
	if err != nil {
		logrus.Errorf("<%v>[service.StartSession] could not get app-config data: %v\n", r.Host, err)
		server.onErr(w, http.StatusInternalServerError, "Could not load meta data")
		return
	}
	server.onSuccess(w, http.StatusOK, map[string]interface{}{
		"ticket": wsTicket,
		"meta":   config,
	})
}
