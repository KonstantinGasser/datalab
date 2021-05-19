package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.notification-live/notifyhub"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) HandleIncomingNofication(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[handler.IncomingNofication] received require\n")
	var notification struct {
		ReceiverUuid string            `json:"receiver_uuid"`
		ReceiverOrgn string            `json:"receiver_orgn"`
		Msg          notifyhub.Message `json:"msg"`
	}
	if err := handler.decode(r.Body, &notification); err != nil {
		logrus.Errorf("[handler.IncomingNotification] could not decode r.Body: %v\n", err)
		handler.onError(w, "could not decode r.Body", http.StatusBadGateway)
		return
	}
	handler.domain.EventAppInvite(r.Context(), notification.Msg, notification.ReceiverUuid, notification.ReceiverOrgn)

}
