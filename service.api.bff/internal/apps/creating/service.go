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
}

type service struct {
	appMetaService client.ClientAppMeta
}

func NewService(appMetaService client.ClientAppMeta) Service {
	return &service{
		appMetaService: appMetaService,
	}
}

func (s service) CreateApp(ctx context.Context, r *apps.CreateAppRequest) *apps.CreateAppResponse {
	if err := required.Atomic(r); err != nil {
		return &apps.CreateAppResponse{
			Stauts: http.StatusOK,
			Msg:    "Mandatory fields missing",
			Err:    err.Error(),
		}
	}
	err := s.appMetaService.CreateApp(ctx, r)
	if err != nil {
		return &apps.CreateAppResponse{
			Stauts: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}

	return &apps.CreateAppResponse{
		Stauts: http.StatusOK,
		Msg:    "Create App",
	}
}
