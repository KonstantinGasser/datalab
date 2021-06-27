package fetching

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	GetById(ctx context.Context, uuid string) (*appconfigs.AppConfig, errors.Api)
}

type service struct {
	repo appconfigs.AppconfigRepo
}

func NewService(repo appconfigs.AppconfigRepo) Service {
	return &service{repo: repo}
}

func (s service) GetById(ctx context.Context, uuid string) (*appconfigs.AppConfig, errors.Api) {
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return nil, errors.New(http.StatusUnauthorized, fmt.Errorf("missing authentication"), "User not authenticated")
	}

	var storedAppConfig appconfigs.AppConfig
	if err := s.repo.GetById(ctx, uuid, &storedAppConfig); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusNotFound, err, "Could not find App Config")
		}
		return nil, errors.New(http.StatusInternalServerError, err, "Could not get App Config")
	}

	if err := storedAppConfig.HasReadOrWrite(authedUser.Uuid); err != nil {
		return nil, errors.New(http.StatusUnauthorized, err, "User hat no permissions to get App Config")
	}

	return &storedAppConfig, nil
}
