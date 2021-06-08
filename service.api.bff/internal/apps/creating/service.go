package creating

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client"
	"github.com/KonstantinGasser/required"
	"github.com/sirupsen/logrus"
)

type Service interface {
	CreateApp(ctx context.Context, r *apps.CreateAppRequest) *apps.CreateAppResponse
	CreateAppToken(ctx context.Context, r *apps.CreateAppTokenRequest) *apps.CreateAppTokenResponse
}

type service struct {
	appMetaService   client.ClientAppMeta
	appTokenService  client.ClientAppToken
	appConfigService client.ClientAppConfig
	notifyLiveClient client.ClientNotifiyLive
}

func NewService(
	appMetaService client.ClientAppMeta,
	appTokenService client.ClientAppToken,
	appConfigService client.ClientAppConfig,
	notifyLiveClient client.ClientNotifiyLive,
) Service {
	return &service{
		appMetaService:   appMetaService,
		appTokenService:  appTokenService,
		appConfigService: appConfigService,
		notifyLiveClient: notifyLiveClient,
	}
}

func (s service) CreateApp(ctx context.Context, r *apps.CreateAppRequest) *apps.CreateAppResponse {
	if err := required.Atomic(r); err != nil {
		return &apps.CreateAppResponse{
			Status: http.StatusOK,
			Msg:    "Mandatory fields missing",
			Err:    err.Error(),
		}
	}
	appUuid, err := s.appMetaService.CreateApp(ctx, r)
	if err != nil {
		return &apps.CreateAppResponse{
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}

	return &apps.CreateAppResponse{
		Status:  http.StatusOK,
		Msg:     "Create App",
		AppUuid: appUuid,
	}
}

func (s service) CreateAppToken(ctx context.Context, r *apps.CreateAppTokenRequest) *apps.CreateAppTokenResponse {

	appAccessToken, err := s.appTokenService.IssueAppToken(ctx, r)
	if err != nil {
		return &apps.CreateAppTokenResponse{
			Status:   err.Code(),
			Msg:      err.Info(),
			Err:      err.Error(),
			AppToken: nil,
		}
	}
	lockAppErr := s.appMetaService.LockApp(ctx, r.AppUuid, r.AuthedUser)
	lockConfigErr := s.appConfigService.LockAppConfig(ctx, r.AppUuid, r.AuthedUser)
	if lockAppErr != nil || lockConfigErr != nil {
		logrus.Errorf("[creating.CreateAppToken] well this is where compensating function would be a cool thing, right??")
		logrus.Errorf("[app lock err] : %v\n[conf lock err] %v\n", lockAppErr, lockConfigErr)
	}
	s.notifyLiveClient.EmitSendEvent(ctx, 2, client.MutationSyncApp, "", r.AuthedUser.Organization, map[string]interface{}{
		"app_uuid": r.AppUuid,
		"sync":     true,
	})
	return &apps.CreateAppTokenResponse{
		Status:   http.StatusOK,
		Msg:      "App Token created",
		AppToken: appAccessToken,
	}
}
