package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.notification-live/notifyhub"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) HandleRemoveNotification(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[handler.RemoveNotification] received require\n")
	var notification notifyhub.RemoveEvent
	if err := handler.decode(r.Body, &notification); err != nil {
		logrus.Errorf("[handler.IncomingNotification] could not decode r.Body: %v\n", err)
		handler.onError(w, "could not decode r.Body", http.StatusBadGateway)
		return
	}
}
