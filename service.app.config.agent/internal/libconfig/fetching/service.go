package fetching

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/libconfig"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	FromStore(ctx context.Context, appUuid string) (*libconfig.Config, errors.Api)
}

type service struct {
	repo libconfig.LibClientRepo
}

func NewService(repo libconfig.LibClientRepo) Service {
	return &service{
		repo: repo,
	}
}

func (s service) FromStore(ctx context.Context, appUuid string) (*libconfig.Config, errors.Api) {
	config, err := s.repo.Load(ctx, appUuid)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusBadRequest, err, "Could not find any related data")
		}
		return nil, errors.New(http.StatusInternalServerError, err, "Could not load config for client-lib")
	}
	return config, nil
}
