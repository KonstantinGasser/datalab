package initializing

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens"
)

type Service interface {
	InitializeAppToken(ctx context.Context, appUuid, appHash, appOwner string) errors.Api
}

type service struct {
	repo apptokens.ApptokenRepo
}

func NewService(repo apptokens.ApptokenRepo) Service {
	return &service{repo: repo}
}

// InitializeAppToken creates the core data object to represent an AppToken and stores it in the
// database
func (s *service) InitializeAppToken(ctx context.Context, appUuid, appHash, appOwner string) errors.Api {
	appToken, err := apptokens.NewDefault(appUuid, appHash, appOwner)
	if err != nil {
		return errors.New(http.StatusBadRequest, err, "Could not create App Token")
	}
	err = s.repo.Initialize(ctx, *appToken)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not create App Token")
	}
	return nil
}
