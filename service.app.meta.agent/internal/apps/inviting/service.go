package inviting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports/client"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	SendInvite(ctx context.Context, appUuid, invitedUuid string) (string, errors.Api)
	SendInviteReminderOK(ctx context.Context, appUuid, userUuid string) errors.Api
	AcceptInvite(ctx context.Context, appUuid string) errors.Api
}

type service struct {
	repo           apps.AppsRepository
	userAuthClient *client.ClientUserAuth
}

func NewService(repo apps.AppsRepository, userAuthClient *client.ClientUserAuth) Service {
	return &service{
		repo:           repo,
		userAuthClient: userAuthClient,
	}
}

// SendInvite adds the given user to the App.Member in state InvitePending
func (s service) SendInvite(ctx context.Context, appUuid, invitedUuid string) (string, errors.Api) {
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return "", errors.New(http.StatusUnauthorized, fmt.Errorf("missing authentication"), "User not authenticated")
	}

	var storedApp apps.App
	if err := s.repo.GetById(ctx, appUuid, &storedApp); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New(http.StatusNotFound,
				err,
				"Could not find App data")
		}
		return "", errors.New(http.StatusInternalServerError,
			err,
			"Could not get App data")
	}
	// only onwer may invite member to app
	if err := storedApp.IsOwner(authedUser.Uuid); err != nil {
		return "", errors.New(http.StatusUnauthorized,
			err,
			"Only Owner can invite members")
	}

	member, inviteErr := storedApp.AddInvite(invitedUuid)
	if inviteErr != nil {
		return "", errors.New(http.StatusBadRequest,
			inviteErr,
			"User is already member of App")
	}
	err := s.repo.AddMember(ctx, appUuid, *member)
	if err != nil {
		return "", errors.New(http.StatusInternalServerError,
			err,
			"Could not add Invite to App")
	}
	return storedApp.Name, nil
}

func (s service) SendInviteReminderOK(ctx context.Context, appUuid, userUuid string) errors.Api {
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
	reminderOK := storedApp.InviteReminderOk(userUuid)
	if !reminderOK {
		return errors.New(http.StatusBadRequest,
			fmt.Errorf("invite reminder cannot be send to user"),
			"Invite reminder cannot be send to User")
	}
	return nil
}

// AcceptInvite updates the App.Member for the given user to state InviteAccepted
func (s service) AcceptInvite(ctx context.Context, appUuid string) errors.Api {
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return errors.New(http.StatusUnauthorized, fmt.Errorf("missing authentication"), "User not authenticated")
	}

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

	// user must have an open invite for the app
	openInvite := storedApp.OpenInvite(authedUser.Uuid)
	if openInvite == nil {
		return errors.New(http.StatusUnauthorized,
			fmt.Errorf("could not find open invite for user"),
			"User does not have an open invite")
	}
	err := s.repo.MemberStatus(ctx, appUuid, *openInvite)
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not update invite status")
	}
	// append users permission with new app
	if err := s.userAuthClient.AddAppAccess(ctx, openInvite.Uuid, storedApp.Uuid); err != nil {
		return err
	}
	return nil
}
