package fetching

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	FetchAppByID(ctx context.Context, appUuid string) (*apps.App, errors.Api)
	FetchAppSubsets(ctx context.Context) ([]apps.AppSubset, errors.Api)
}

type service struct {
	repo apps.AppsRepository
}

func NewService(repo apps.AppsRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) FetchAppByID(ctx context.Context, appUuid string) (*apps.App, errors.Api) {
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return nil, errors.New(http.StatusUnauthorized, fmt.Errorf("missing authentication"), "User not authenticated")
	}

	var app apps.App
	err := s.repo.GetById(ctx, appUuid, &app)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusNotFound, err, "Could not find App data")
		}
		return nil, errors.New(http.StatusInternalServerError, err, "Could not get App data")
	}
	// check for permissions
	if err := app.HasReadAccess(authedUser.Uuid); err != nil {
		return nil, errors.New(http.StatusUnauthorized, err, "User has no permissions to read App")
	}
	return &app, nil
}

// FetchAllApps collects all Apps the user has permission to read from
func (s service) FetchAppSubsets(ctx context.Context) ([]apps.AppSubset, errors.Api) {
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return nil, errors.New(http.StatusUnauthorized, fmt.Errorf("missing authentication"), "User not authenticated")
	}

	var storedApps []apps.App
	err := s.repo.GetAll(ctx, authedUser.Uuid, &storedApps)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err, "Could not get Apps")
	}
	return apps.SubsetOf(storedApps...), nil
}
