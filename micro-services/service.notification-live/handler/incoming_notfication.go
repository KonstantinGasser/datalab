package handler

import (
	"net/http"
	"time"

	"github.com/KonstantinGasser/datalab/service.notification-live/notifyhub"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) HandleIncomingNotification(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[handler.IncomingNotification] received require\n")
	var notification notifyhub.IncomingEvent
	if err := handler.decode(r.Body, &notification); err != nil {
		logrus.Errorf("[handler.IncomingNotification] could not decode r.Body: %v\n", err)
		handler.onError(w, "could not decode r.Body", http.StatusBadGateway)
		return
	}
	notification.Timestamp = time.Now().Unix()

	handler.domain.EventAppInvite(r.Context(), notification)
}
