package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.notification-live/notifyhub"
	userauthsvc "github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
)

type NotificationLogic interface {
	IsLoggedIn(ctx context.Context, token string) (*common.AuthedUser, errors.Api)
	OpenSocket(ctx context.Context, w http.ResponseWriter, r *http.Request) error
	EventAppInvite(ctx context.Context, notification notifyhub.IncomingEvent)

	HideNotification(ctx context.Context, notify notifyhub.HideEvent)
	RemoveNotifcation(ctx context.Context, notify notifyhub.RemoveEvent)
}

type notificationlogic struct {
	userauthClient userauthsvc.UserAuthenticationClient
	notifyHub      *notifyhub.NotifyHub
}

func NewNotificationLogic(userauthC userauthsvc.UserAuthenticationClient, notifyHub *notifyhub.NotifyHub) NotificationLogic {
	return &notificationlogic{
		userauthClient: userauthC,
		notifyHub:      notifyHub,
	}
}
