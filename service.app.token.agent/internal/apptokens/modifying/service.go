package modifying

import (
	"context"
	"net/http"
	"time"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	IssueAppToken(ctx context.Context, appUuid, callerUuid string) (string, time.Time, errors.Api)
}

type service struct {
	repo apptokens.ApptokenRepo
}

func NewService(repo apptokens.ApptokenRepo) Service {
	return &service{repo: repo}
}

// IssueAppToken issues an new Jwt based on the stored App Token information and updates its
// expiration time. Returns the new Jwt and Exp
func (s *service) IssueAppToken(ctx context.Context, appUuid, callerUuid string) (string, time.Time, errors.Api) {
	var storedAppToken apptokens.AppToken
	if err := s.repo.GetById(ctx, appUuid, &storedAppToken); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", time.Time{}, errors.New(http.StatusNotFound, err, "Could not find App Token data")
		}
		return "", time.Time{}, errors.New(http.StatusInternalServerError, err, "Could not get App Token data")
	}
	// verify that the user is the actual owner of the AppToken
	if err := storedAppToken.HasReadWrite(callerUuid); err != nil {
		return "", time.Time{}, errors.New(http.StatusUnauthorized, err, "User must be owner to generate AppToken")
	}

	modifiedAppToken, err := storedAppToken.Issue()
	if err != nil {
		if err == apptokens.ErrAppTokenStillValid {
			return "", time.Time{}, errors.New(http.StatusBadRequest, err, "Current App Token is still valid")
		}
		return "", time.Time{}, errors.New(http.StatusInternalServerError, err, "Could not issue new App Token")
	}
	repoErr := s.repo.Update(ctx, modifiedAppToken.AppRefUuid, modifiedAppToken.Jwt, modifiedAppToken.Exp)
	if repoErr != nil {
		return "", time.Time{}, errors.New(http.StatusInternalServerError, repoErr, "Could not update App Token")
	}
	return modifiedAppToken.Jwt, modifiedAppToken.Exp, nil
}
