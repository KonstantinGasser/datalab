package inviting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	SendInvite(ctx context.Context, appUuid, invitedUuid string) (string, errors.Api)
	SendInviteReminderOK(ctx context.Context, appUuid, userUuid string) errors.Api
	AcceptInvite(ctx context.Context, appUuid string) errors.Api
}

type service struct {
	repo            apps.AppsRepository
	emitterAppConf  ports.EventEmitter
	emitterAppToken ports.EventEmitter
	// userAuthClient *client.ClientUserAuth
}

func NewService(repo apps.AppsRepository, emitterAppConfSvc ports.EventEmitter, emitterAppTknSvc ports.EventEmitter) Service {
	return &service{
		repo:            repo,
		emitterAppConf:  emitterAppConfSvc,
		emitterAppToken: emitterAppTknSvc,
		// userAuthClient: userAuthClient,
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
	err := s.repo.MemberStatus(ctx, appUuid, *openInvite, apps.InviteAccepted)
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not update invite status")
	}

	// tell other app related services to add new user as member of app
	if err := s.emitAppendPermissions(ctx, appUuid, authedUser.Uuid); err != nil {
		// if either call to a service failes - rollback all
		if err := s.compensateInvite(ctx, appUuid, apps.Member{
			Uuid:   authedUser.Uuid,
			Status: apps.InvitePending,
		}); err != nil {
			logrus.Errorf("[invite.Rollback] could not rollback App Invite accept: %v\n", err)
		}
		if err := s.emitRollbackAppendPermissions(ctx, appUuid, authedUser.Uuid); err != nil {
			logrus.Errorf("[invite.Rollback] could not rollback append permissions: %v\n", err)
		}
		return errors.New(http.StatusInternalServerError, err, "Could not accept App-Invite")
	}
	return nil
}

// emitInitEvent distributes the event that a new app has been created triggering the init endpoints
// of the AppTokenService and AppConfigService
func (s service) emitAppendPermissions(ctx context.Context, appUuid, joinedUser string) error {
	withCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	var errC = make(chan error)
	emitterEvent := ports.NewPermissionEvent(appUuid, joinedUser)

	go s.emitterAppToken.EmitAppendPermissions(withCancel, emitterEvent, errC)
	go s.emitterAppConf.EmitAppendPermissions(withCancel, emitterEvent, errC)
	// go s.emitterUserPermissions.Emit(withCancel, emitterEvent, errC)

	for i := 0; i < 2; i++ {
		err := <-errC
		if err != nil {
			logrus.Errorf("[%s][inviting.EmitAppendPermissions] emit cause error: %v\n", ctx.Value("tracingID"), err)
			// if there is an error while emitting events
			// here the emiited events must succed in order for the
			// transaction to succeed - hence if err cancel context and
			// role back (if that would have been implmeneted)
			return err
		}
	}
	close(errC)
	return nil
}

// emitInitEvent distributes the event that a new app has been created triggering the init endpoints
// of the AppTokenService and AppConfigService
func (s service) emitRollbackAppendPermissions(ctx context.Context, appUuid, joinedUser string) error {
	withCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	var errC = make(chan error)
	emitterEvent := ports.NewPermissionEvent(appUuid, joinedUser)

	go s.emitterAppToken.EmitRollbackAppendPermissions(withCancel, emitterEvent, errC)
	go s.emitterAppConf.EmitRollbackAppendPermissions(withCancel, emitterEvent, errC)

	// go s.emitterUserPermissions.Emit(withCancel, emitterEvent, errC)

	for i := 0; i < 2; i++ {
		err := <-errC
		if err != nil {
			logrus.Errorf("[%s][inviting.EmitRollbackAppendPermissions] emit cause error: %v\n", ctx.Value("tracingID"), err)
			// if there is an error while emitting the rollback event
			// ignore error since at least one error is to be expected since an error happend in the forward
			// transaction
			continue
		}
	}
	close(errC)
	return nil
}
