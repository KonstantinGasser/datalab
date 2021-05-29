package inviting

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client"
	"github.com/KonstantinGasser/required"
	"github.com/sirupsen/logrus"
)

type Service interface {
	SendInvite(ctx context.Context, r *apps.SendInviteRequest) *apps.SendInviteResponse
	AcceptInvite(ctx context.Context, r *apps.AcceptInviteRequest) *apps.AcceptInviteResponse
}

type service struct {
	appMetaClient    client.ClientAppMeta
	notifyLiveClient client.ClientNotifiyLive
}

func NewService(appMetaClient client.ClientAppMeta, notifyLiveClient client.ClientNotifiyLive) Service {
	return &service{
		appMetaClient:    appMetaClient,
		notifyLiveClient: notifyLiveClient,
	}
}

func (s service) SendInvite(ctx context.Context, r *apps.SendInviteRequest) *apps.SendInviteResponse {
	if err := required.Atomic(r); err != nil {
		return &apps.SendInviteResponse{
			Status: http.StatusBadRequest,
			Msg:    "Mandatory fields missing",
			Err:    err.Error(),
		}
	}

	appName, err := s.appMetaClient.SendInvite(ctx, r)
	if err != nil {
		return &apps.SendInviteResponse{
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}
	notifyErr := s.notifyLiveClient.EmitSendInvite(ctx, r.InvitedUuid, r.AuthedUser.Organization, map[string]interface{}{
		"app_uuid":  r.AppUuid,
		"app_name":  appName,
		"app_owner": r.AuthedUser.Username,
	})
	// if message not send for now I dont care...will change in future
	if notifyErr != nil {
		logrus.Errorf("[service.inviting.SendInvite] could not send invite to notification service: %v\n", notifyErr)
	}
	return &apps.SendInviteResponse{
		Status: http.StatusOK,
		Msg:    "Invite send",
	}
}

func (s service) AcceptInvite(ctx context.Context, r *apps.AcceptInviteRequest) *apps.AcceptInviteResponse {
	if err := required.Atomic(r); err != nil {
		return &apps.AcceptInviteResponse{
			Status: http.StatusBadRequest,
			Msg:    "Mandatory fields missing",
			Err:    err.Error(),
		}
	}

	err := s.appMetaClient.AcceptInvite(ctx, r)
	if err != nil {
		return &apps.AcceptInviteResponse{
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}
	return &apps.AcceptInviteResponse{
		Status: http.StatusOK,
		Msg:    "Accepted App Invite",
	}
}
