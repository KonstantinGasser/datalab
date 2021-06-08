package updating

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
	LockApp(ctx context.Context, appUuid string, authedUser *common.AuthedUser) errors.Api
}

type service struct {
	repo apps.AppsRepository
}

func NewService(repo apps.AppsRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) LockApp(ctx context.Context, appUuid string, authedUser *common.AuthedUser) errors.Api {
	var storedApp apps.App
	if err := s.repo.GetById(ctx, appUuid, &storedApp); err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New(http.StatusNotFound,
				err,
				"Could not find App data")
		}
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not get App data")
	}

	if storedApp.IsLocked() {
		return errors.New(http.StatusBadRequest,
			fmt.Errorf("app already locked"),
			"App is already locked")
	}

	if err := storedApp.IsOwner(authedUser.Uuid); err != nil {
		return errors.New(http.StatusUnauthorized,
			err,
			"User not authorized to perform action")
	}

	storedApp.Locked = true
	if err := s.repo.SetAppLock(ctx, storedApp.Uuid); err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New(http.StatusNotFound,
				err,
				"Could not find App data")
		}
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not update App data")
	}
	return nil
}
