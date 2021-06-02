package domain

import (
	"context"

	"github.com/KonstantinGasser/datalab/service.notification-live/notifyhub"
)

func (svc notificationlogic) RemoveNotifcation(ctx context.Context, notify notifyhub.RemoveEvent) {
	svc.notifyHub.RemoveNotify <- &notify
}

func (svc notificationlogic) HideNotification(ctx context.Context, notify notifyhub.HideEvent) {
	svc.notifyHub.HideNotify <- &notify
}
