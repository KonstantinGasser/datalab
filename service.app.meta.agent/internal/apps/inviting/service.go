package inviting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps"
	"github.com/KonstatinGasser/datalab/service.app.meta.agent/internal/apps"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	SendInvite(ctx context.Context, appUuid, invitedUuid string, authedUser *common.AuthedUser) errors.Api
	AcceptInvite(ctx context.Context, appUuid, userUuid string) errors.Api
}

type service struct {
	repo apps.AppsRepository
}

func NewService(repo apps.AppsRepository) Service {
	return &service{repo: repo}
}

// SendInvite adds the given user to the App.Member in state InvitePending
func (s service) SendInvite(ctx context.Context, appUuid, invitedUuid string, authedUser *common.AuthedUser) errors.Api {
	var storedApp apps.App
	if err := s.repo.GetById(ctx, appUuid, &storedApp); err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New(http.StatusNotFound, err, "Could not find App data")
		}
		return errors.New(http.StatusInternalServerError, err, "Could not get App data")
	}
	// only onwer may invite member to app
	if err := storedApp.IsOwner(authedUser.Uuid); err != nil {
		return errors.New(http.StatusUnauthorized, err, "Only Owner can invite members")
	}

	member, inviteErr := storedApp.AddInvite(invitedUuid)
	if inviteErr != nil {
		return errors.New(http.StatusBadRequest, inviteErr, "User is already member of App")
	}
	err := s.repo.AddMember(ctx, appUuid, *member)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not add Invite to App")
	}
	return nil
}

// AcceptInvite updates the App.Member for the given user to state InviteAccepted
func (s service) AcceptInvite(ctx context.Context, appUuid, userUuid string) errors.Api {
	var storedApp apps.App
	if err := s.repo.GetById(ctx, appUuid, &storedApp); err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New(http.StatusNotFound, err, "Could not find App data")
		}
		return errors.New(http.StatusInternalServerError, err, "Could not get App data")
	}

	// user must have an open invite for the app
	openInvite := storedApp.OpenInvite(userUuid)
	if openInvite == nil {
		return errors.New(http.StatusUnauthorized, fmt.Errorf("could not find open invite for user"), "User does not have an open invite")
	}
	err := s.repo.MemberStatus(ctx, appUuid, *openInvite)
	if err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not update invite status")
	}
	return nil
}
