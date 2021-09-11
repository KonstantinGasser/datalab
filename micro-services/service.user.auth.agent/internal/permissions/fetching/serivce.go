package fetching

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	GetById(ctx context.Context, userUuid string) (*permissions.Permission, errors.Api)
}

type service struct {
	repo permissions.PermissionRepo
}

func NewService(repo permissions.PermissionRepo) Service {
	return &service{
		repo: repo,
	}
}

func (s service) GetById(ctx context.Context, userUuid string) (*permissions.Permission, errors.Api) {

	var storedPermission permissions.Permission
	err := s.repo.GetById(ctx, userUuid, &storedPermission)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusNotFound,
				err,
				"Could not find User Permission")
		}
		return nil, errors.New(http.StatusInternalServerError, err, "Could not get User Permission")
	}

	return &storedPermission, nil
}
