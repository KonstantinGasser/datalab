package fetching

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.eventmanager.live/ports/client"
)

type Service interface {
	FetchConfigMetaData(ctx context.Context, appUuid string) (*common.AppConfigurations, errors.Api)
}

type service struct {
	appConfigClient client.ClientAppConfig
}

func NewService(appConfigClient client.ClientAppConfig) Service {
	return &service{appConfigClient}
}

// ValidateAppToken issues an new Jwt based on the stored App Token information and updates its
// expiration time. Returns the new Jwt and Exp
func (s *service) FetchConfigMetaData(ctx context.Context, appUuid string) (*common.AppConfigurations, errors.Api) {

	config, err := s.appConfigClient.GetAppConfig(ctx, appUuid, nil)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err, "could not load meta data")
	}
	return config, nil
}
