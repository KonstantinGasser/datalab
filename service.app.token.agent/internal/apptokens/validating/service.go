package validating

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	ValidateAppToken(ctx context.Context, jwt string) (string, string, errors.Api)
	InvalidateAppToken(ctx context.Context, jwt string) errors.Api
}

type service struct {
	repo apptokens.ApptokenRepo
}

func NewService(repo apptokens.ApptokenRepo) Service {
	return &service{repo: repo}
}

// ValidateAppToken issues an new Jwt based on the stored App Token information and updates its
// expiration time. Returns the new Jwt and Exp
func (s *service) ValidateAppToken(ctx context.Context, jwt string) (string, string, errors.Api) {

	appUuid, appOrigin, refreshCount, err := apptokens.ClaimsFromJwt(jwt)
	if err != nil {
		return "", "", errors.New(http.StatusUnauthorized,
			err,
			"App Token corrputed")
	}

	var storedAppToken apptokens.AppToken
	if err := s.repo.GetById(ctx, appUuid, &storedAppToken); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", "", errors.New(http.StatusNotFound,
				err,
				"Could not find App Token data")
		}
		return "", "", errors.New(http.StatusInternalServerError,
			err,
			"Could not get App Token data")
	}

	if !storedAppToken.IsValid(refreshCount) {
		return "", "", errors.New(http.StatusUnauthorized,
			fmt.Errorf("app-token is no longer valid"), "This App-Token has been invalidated")
	}
	return appUuid, appOrigin, nil
}

func (s *service) InvalidateAppToken(ctx context.Context, appUuid string) errors.Api {
	var stored apptokens.AppToken
	if err := s.repo.GetById(ctx, appUuid, &stored); err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New(http.StatusBadRequest,
				err,
				"Could not find App Token data")
		}
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not get App Token data")
	}
	return nil
}
