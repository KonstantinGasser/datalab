package initializing

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs"
)

type Service interface {
	InitializeAppConfig(ctx context.Context, appRefUuid, appOwner, ownerOrgn string, isPrivate bool) errors.Api
}

type service struct {
	repo appconfigs.AppconfigRepo
}

func NewService(repo appconfigs.AppconfigRepo) Service {
	return &service{repo: repo}
}

// InitializeAppConfig creates the core data object to represent an AppConfig and stores it in the
// database
func (s *service) InitializeAppConfig(ctx context.Context, appRefUuid, appOwner, ownerOrgn string, isPrivate bool) errors.Api {
	appConfig := appconfigs.NewDefault(appRefUuid, appOwner, ownerOrgn, isPrivate)
	err := s.repo.Initialize(ctx, *appConfig)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not create default App Config")
	}
	return nil
}
