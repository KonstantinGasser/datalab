package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (handler *Handler) HandlerOpenSocket(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[service.eventmanager-live.OpenSocket] received request\n", r.Host)

	ticket, ok := r.Context().Value(typeKeyTicket(keyTicket)).(string)
	if !ok {
		handler.onErr(w, http.StatusUnauthorized, "web-socket ticket not present")
		return
	}
	if ticket == "" {
		handler.onErr(w, http.StatusUnauthorized, "web-socket ticket not present")
		return
	}

	err := handler.domain.OpenSocket(r.Context(), ticket, w, r)
	if err != nil {
		logrus.Errorf("<%v>[service.eventmanager-live.OpenWebSocket] could not establish connection: %v\n", r.Host, err.Error())
		handler.onErr(w, int(err.Code()), err.Info())
		return
	}
}
