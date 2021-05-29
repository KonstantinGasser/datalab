package creating

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client"
	"github.com/KonstantinGasser/required"
)

type Service interface {
	CreateApp(ctx context.Context, r *apps.CreateAppRequest) *apps.CreateAppResponse
	CreateAppToken(ctx context.Context, r *apps.CreateAppTokenRequest) *apps.CreateAppTokenResponse
}

type service struct {
	appMetaService  client.ClientAppMeta
	appTokenService client.ClientAppToken
}

func NewService(appMetaService client.ClientAppMeta, appTokenService client.ClientAppToken) Service {
	return &service{
		appMetaService:  appMetaService,
		appTokenService: appTokenService,
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
	return &apps.CreateAppTokenResponse{
		Status:   http.StatusOK,
		Msg:      "App Token created",
		AppToken: appAccessToken,
	}
}
