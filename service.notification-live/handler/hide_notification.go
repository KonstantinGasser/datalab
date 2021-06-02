package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.notification-live/notifyhub"
	"github.com/sirupsen/logrus"
)

func (handler Handler) HandleHideNotification(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[handler.HideNotification] received require\n")
	var notification notifyhub.HideEvent
	if err := handler.decode(r.Body, &notification); err != nil {
		logrus.Errorf("[handler.HideNotification] could not decode r.Body: %v\n", err)
		handler.onError(w, "could not decode r.Body", http.StatusBadGateway)
		return
	}
	handler.domain.HideNotification(r.Context(), notification)
}
