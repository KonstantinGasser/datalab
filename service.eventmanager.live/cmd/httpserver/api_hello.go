package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.eventmanager-live/jwts"
	"github.com/KonstantinGasser/datalab/service.eventmanager.live/internal/sessions"
	"github.com/sirupsen/logrus"
)

func (server *Server) Hello(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[server.Hello] received request\n")

	var session sessions.Session
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		server.onErr(w, http.StatusBadRequest, "could not decode r.Body")
		return
	}

	appUuid := ctx_value.GetString(r.Context(), "app.uuid")
	if appUuid == "" {
		server.onErr(w, http.StatusUnauthorized, "not authorized")
		return
	}
	appOrigin := ctx_value.GetString(r.Context(), "app.origin")
	if appOrigin == "" {
		server.onErr(w, http.StatusUnauthorized, "not authorized")
		return
	}
	var request = sessions.StartRequest{
		AppUuid:   appUuid,
		AppOrigin: appOrigin,
		Session:   &session,
	}
	config, err := server.sessionSvc.FetchConfigMetaData(r.Context(), request.AppUuid)
	if err != nil {
		server.onErr(w, http.StatusInternalServerError, "could not load meta data")
		return
	}
	cookie, ok := r.Context().Value(typeKeyCookie(keyCookie)).(*http.Cookie)
	if !ok || cookie == nil {
		logrus.Errorf("<%v>[service.StartSession] could not find cookie\n", r.Host)
		server.onErr(w, http.StatusInternalServerError, "missing credentials")
		return
	}
	wsTicket, errTicket := jwts.WebSocketTicket(cookie.Name)
	if errTicket != nil {
		logrus.Errorf("<%v>[service.StartSession] could not issue ws ticket\n", r.Host)
		server.onErr(w, http.StatusInternalServerError, "could not issue ws ticket")
		return
	}
	server.onSuccess(w, http.StatusOK, map[string]interface{}{
		"ticket":   wsTicket,
		"btn_defs": config.GetBtnTime(),
		"stages":   config.GetFunnel(),
	})
}
