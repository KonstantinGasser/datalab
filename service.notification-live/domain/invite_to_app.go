package domain

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.notification-live/notifyhub"
)

func (svc notificationlogic) EventAppInvite(ctx context.Context, notification notifyhub.IncomingEvent) {
	svc.notifyHub.Notify <- &notification
}
