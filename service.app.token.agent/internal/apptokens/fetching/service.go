package fetching

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens"
)

type Service interface {
	FetchById(ctx context.Context, appUuid string) (*apptokens.AppToken, errors.Api)
}

type service struct {
	repo apptokens.ApptokenRepo
}

func NewService(repo apptokens.ApptokenRepo) Service {
	return &service{repo: repo}
}

// FetchById returns the stores AppToken if the user has read permissions on the AppToken
func (s *service) FetchById(ctx context.Context, appUuid string) (*apptokens.AppToken, errors.Api) {
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return nil, errors.New(http.StatusUnauthorized, fmt.Errorf("missing authentication"), "User not authenticated")
	}

	var storedAppToken apptokens.AppToken
	if err := s.repo.GetById(ctx, appUuid, &storedAppToken); err != nil {
		return nil, errors.New(http.StatusInternalServerError, err, "Could not get AppToken")
	}

	// check if user is allowed to read or write AppToken
	if err := storedAppToken.HasReadOrWrite(authedUser.Uuid, authedUser.ReadWriteApps...); err != nil {
		return nil, errors.New(http.StatusUnauthorized, err, "User is not allowed to access App Token")
	}
	return &storedAppToken, nil
}
