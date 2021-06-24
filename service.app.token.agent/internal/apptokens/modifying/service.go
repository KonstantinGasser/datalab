package modifying

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	IssueAppToken(ctx context.Context, orgn, appName, appUuid string) (string, int64, errors.Api)
	UnlockAppToken(ctx context.Context, appUuid string) errors.Api
}

type service struct {
	repo apptokens.ApptokenRepo
}

func NewService(repo apptokens.ApptokenRepo) Service {
	return &service{repo: repo}
}

// IssueAppToken issues an new Jwt based on the stored App Token information and updates its
// expiration time. Returns the new Jwt and Exp
func (s *service) IssueAppToken(ctx context.Context, orgn, appName, appUuid string) (string, int64, errors.Api) {
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return "", 0, errors.New(http.StatusUnauthorized, fmt.Errorf("missing authentication"), "User not authenticated")
	}

	var storedAppToken apptokens.AppToken
	if err := s.repo.GetById(ctx, appUuid, &storedAppToken); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", 0, errors.New(http.StatusNotFound,
				err,
				"Could not find App Token data")
		}
		return "", 0, errors.New(http.StatusInternalServerError,
			err,
			"Could not get App Token data")
	}
	// verify that the user is the actual owner of the AppToken
	if err := storedAppToken.HasReadWrite(authedUser.Uuid); err != nil {
		return "", 0, errors.New(http.StatusUnauthorized,
			err,
			"User must be owner to generate AppToken")
	}
	if storedAppToken.Locked {
		return "", 0, errors.New(http.StatusUnauthorized,
			fmt.Errorf("app is locked"),
			"App is in locked state - change not possible")
	}

	modifiedAppToken, err := storedAppToken.Issue(orgn, appName)
	if err != nil {
		if err == apptokens.ErrWrongAppHash {
			return "", 0, errors.New(http.StatusUnauthorized,
				err,
				"Organization-Name/App-Name ist incorrect")
		}
		if err == apptokens.ErrAppTokenStillValid {
			return "", 0, errors.New(http.StatusBadRequest,
				err,
				"Current App Token is still valid")
		}
		return "", 0, errors.New(http.StatusInternalServerError,
			err,
			"Could not issue new App Token")
	}

	// update new app-token in database
	repoErr := s.repo.Update(ctx, modifiedAppToken.AppRefUuid, modifiedAppToken.Jwt, modifiedAppToken.Exp)
	if repoErr != nil {
		return "", 0, errors.New(http.StatusInternalServerError,
			repoErr,
			"Could not update App Token")
	}
	return modifiedAppToken.Jwt, modifiedAppToken.Exp, nil
}

// UnlockAppToken deletes the current app token incrementing the refresh-count in order to invalidate the current
// app token
func (s service) UnlockAppToken(ctx context.Context, appUuid string) errors.Api {
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return errors.New(http.StatusUnauthorized, fmt.Errorf("missing authentication"), "User not authenticated")
	}

	var storedAppToken apptokens.AppToken
	if err := s.repo.GetById(ctx, appUuid, &storedAppToken); err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not get AppToken")
	}

	if err := storedAppToken.HasReadWrite(authedUser.Uuid); err != nil {
		return errors.New(http.StatusUnauthorized,
			err,
			"User has no permissions for this action")
	}

	if err := s.repo.SetAppTokenLock(ctx, appUuid, false); err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not unlock App-Token")
	}

	// if an app gets unlocked the current apptoken will get invalid
	// TODO: implement invalidation of app token with cound in JWT and db.
	// for verification app token cound and db count must match else token has been
	// invalidated
	repoErr := s.repo.Update(ctx, appUuid, "", 0)
	if repoErr != nil {
		return errors.New(http.StatusInternalServerError,
			repoErr,
			"Could not reset App Token")
	}
	return nil
}
