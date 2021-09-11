package adding

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions"
)

type Service interface {
	AddApp(ctx context.Context, userUuid, appUuid string) errors.Api
}

type service struct {
	repo permissions.PermissionRepo
}

func NewService(repo permissions.PermissionRepo) Service {
	return &service{
		repo: repo,
	}
}

func (s service) AddApp(ctx context.Context, userUuid, appUuid string) errors.Api {

	if err := s.repo.AddApp(ctx, userUuid, appUuid); err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not add App Permission")
	}
	return nil
}
