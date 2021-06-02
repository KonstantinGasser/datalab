package modifying

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client"
)

type Service interface {
	UpdateConfig(ctx context.Context, r *apps.UpdateConfigRequest) *apps.UpdateConfigResponse
}

type service struct {
	appConfigClient  client.ClientAppConfig
	notifyLiveClient client.ClientNotifiyLive
}

func NewService(appConfigClient client.ClientAppConfig, notifyLiveClient client.ClientNotifiyLive) Service {
	return &service{
		appConfigClient:  appConfigClient,
		notifyLiveClient: notifyLiveClient,
	}
}

func (s service) UpdateConfig(ctx context.Context, r *apps.UpdateConfigRequest) *apps.UpdateConfigResponse {

	err := s.appConfigClient.UpdateConfig(ctx, r)
	if err != nil {
		return &apps.UpdateConfigResponse{
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}
	s.notifyLiveClient.EmitSendEvent(ctx, 2, client.MutationSyncApp, "", r.AuthedUser.Organization, map[string]interface{}{
		"app_uuid": r.AppRefUuid,
		"sync":     true,
	})
	return &apps.UpdateConfigResponse{
		Status: http.StatusOK,
		Msg:    "Updated App Config",
	}
}
