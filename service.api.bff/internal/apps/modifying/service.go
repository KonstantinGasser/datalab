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
	appConfigClient client.ClientAppConfig
}

func NewService(appConfigClient client.ClientAppConfig) Service {
	return &service{
		appConfigClient: appConfigClient,
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
	return &apps.UpdateConfigResponse{
		Status: http.StatusOK,
		Msg:    "Updated App Config",
	}
}
