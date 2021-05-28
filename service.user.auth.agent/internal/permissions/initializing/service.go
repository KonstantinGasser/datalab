package initializing

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions"
)

type Service interface {
	InitPermissions(ctx context.Context, userUuid, userOrgn string) errors.Api
}

type service struct {
	repo permissions.PermissionRepo
}

func NewService(repo permissions.PermissionRepo) Service {
	return &service{
		repo: repo,
	}
}

func (s service) InitPermissions(ctx context.Context, userUuid, userOrgn string) errors.Api {
	permission, err := permissions.NewDefaultPermission(userUuid, userOrgn)
	if err != nil {
		return errors.New(http.StatusBadRequest,
			err,
			"Missing fields")
	}
	err = s.repo.Store(ctx, *permission)
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not initialize User Permission")
	}
	return nil
}
