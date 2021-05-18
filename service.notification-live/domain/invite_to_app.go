package domain

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.notification-live/notifyhub"
)

func (svc notificationlogic) EventAppInvite(ctx context.Context, msg notifyhub.Message, recUuid, recOrgn string) {

	svc.notifyHub.Notify <- &notifyhub.Notification{
		Uuid:         recUuid,
		Organization: recOrgn,
		Msg:          msg,
	}
}
