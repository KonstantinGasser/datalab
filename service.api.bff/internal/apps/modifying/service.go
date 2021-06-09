package modifying

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client"
	"github.com/sirupsen/logrus"
)

type Service interface {
	UpdateConfig(ctx context.Context, r *apps.UpdateConfigRequest) *apps.UpdateConfigResponse
	UnlockApp(ctx context.Context, r *apps.UnlockRequest) *apps.UnlockResponse
}

type service struct {
	appConfigClient  client.ClientAppConfig
	appMetaClient    client.ClientAppMeta
	appTokenClient   client.ClientAppToken
	notifyLiveClient client.ClientNotifiyLive
}

func NewService(
	appConfigClient client.ClientAppConfig,
	appMetaClient client.ClientAppMeta,
	appTokenClient client.ClientAppToken,
	notifyLiveClient client.ClientNotifiyLive) Service {
	return &service{
		appConfigClient:  appConfigClient,
		appMetaClient:    appMetaClient,
		appTokenClient:   appTokenClient,
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

func (s service) UnlockApp(ctx context.Context, r *apps.UnlockRequest) *apps.UnlockResponse {

	appMetaErr := s.appMetaClient.UnlockApp(ctx, r)
	appConfErr := s.appConfigClient.UnlockAppConfig(ctx, r)
	appTokenErr := s.appTokenClient.UnlockAppToken(ctx, r)

	if appMetaErr != nil || appConfErr != nil || appTokenErr != nil {
		logrus.Errorf("[modifying.UnlockApp] could not unlock app: appmeta: %v\nappconfig: %v\napptoken: %v\n", appMetaErr, appConfErr, appTokenErr)
	}

	s.notifyLiveClient.EmitSendEvent(ctx, 2, client.MutationSyncApp, "", r.AuthedUser.Organization, map[string]interface{}{
		"app_uuid": r.AppUuid,
		"sync":     true,
	})
	return &apps.UnlockResponse{
		Status: http.StatusOK,
		Msg:    "App Data unlocked",
	}
}
