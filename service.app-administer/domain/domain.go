package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/create"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/delete"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/get"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/invite"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/token"
	"github.com/KonstantinGasser/datalab/service.app-administer/errors"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	cfgsvc "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	usersvc "github.com/KonstantinGasser/datalab/service.user-administer/proto"
	userauthsvc "github.com/KonstantinGasser/datalab/service.user-authentication/proto"
)

// AppAdmin is the interface for this service implemented all the service logic
type AppAdmin interface {
	Create(ctx context.Context, in *proto.CreateRequest) (string, errors.ErrApi)
	Delete(ctx context.Context, in *proto.DeleteRequest) errors.ErrApi
	GetSingle(ctx context.Context, in *proto.GetRequest) (*common.AppInfo, errors.ErrApi)
	GetMultiple(ctx context.Context, in *proto.GetListRequest) ([]*common.AppMetaInfo, errors.ErrApi)
	MayAcquireToken(ctx context.Context, in *proto.MayAcquireTokenRequest) errors.ErrApi
	InviteToApp(ctx context.Context, in *proto.InviteRequest) (string, string, errors.ErrApi)
}

type appadmin struct {
	userSvc     usersvc.UserAdministerClient
	configSvc   cfgsvc.AppConfigurationClient
	userauthSvc userauthsvc.UserAuthenticationClient
	repo        repo.Repo
}

func NewAppLogic(repo repo.Repo,
	user usersvc.UserAdministerClient, config cfgsvc.AppConfigurationClient,
	userauth userauthsvc.UserAuthenticationClient) AppAdmin {
	return &appadmin{
		repo:        repo,
		userSvc:     user,
		configSvc:   config,
		userauthSvc: userauth,
	}
}

// Create handles the creation of a new app coordinating the initialization of the app-config
// and performs a rollback if initialization fails
func (svc appadmin) Create(ctx context.Context, in *proto.CreateRequest) (string, errors.ErrApi) {
	appUuid, err := create.App(ctx, svc.repo, in)
	if err != nil {
		if err == create.ErrAppNameExists {
			return "", errors.ErrAPI{
				Status: http.StatusBadRequest,
				Msg:    "App Name already used",
				Err:    err,
			}
		}
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create new App",
			Err:    err,
		}
	}
	// forward uuid of app to app-config service in order for it
	// to create a record to store app-configurations
	respCfg, err := svc.configSvc.Init(ctx, &cfgsvc.InitRequest{
		Tracing_ID: in.GetTracing_ID(),
		ForApp:     appUuid,
	})
	if err != nil || respCfg.GetStatusCode() != http.StatusOK {
		// if the init of the app-config service fails the created app must be deleted
		// to avoid an inconsistent state of the system
		if err := create.CompensateApp(ctx, svc.repo, appUuid); err != nil {
			return "", errors.ErrAPI{
				Status: http.StatusInternalServerError,
				Msg:    "Could not rollback creation of App",
				Err:    err,
			}
		}
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create new App",
			Err:    err,
		}
	}
	respPermissions, err := svc.userauthSvc.AddAppAccess(ctx, &userauthsvc.AddAppAccessRequest{
		Tracing_ID: in.GetTracing_ID(),
		UserUuid:   in.GetOwnerUuid(),
		AppUuid:    appUuid,
		AppRole:    common.AppRole_OWNER,
	})
	if err != nil || respPermissions.GetStatusCode() != http.StatusOK {
		// if the init of the app-config service fails the created app must be deleted
		// to avoid an inconsistent state of the system
		if err := create.CompensateApp(ctx, svc.repo, appUuid); err != nil {
			return "", errors.ErrAPI{
				Status: http.StatusInternalServerError,
				Msg:    "Could not rollback creation of App",
				Err:    err,
			}
		}
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create new App",
			Err:    fmt.Errorf("could not compensate created app and role back: %v", err),
		}
	}
	return appUuid, nil
}

// Delete removes an existing App-Record from the database
func (svc appadmin) Delete(ctx context.Context, in *proto.DeleteRequest) errors.ErrApi {

	_, err := delete.App(ctx, svc.repo, in)
	if err != nil {
		if err == delete.ErrNoPermissions {
			return errors.ErrAPI{Status: http.StatusUnauthorized, Msg: err.Error(), Err: err}
		}
		return errors.ErrAPI{Status: http.StatusInternalServerError, Msg: "Could not delete App", Err: err}
	}

	// here goes the request to delete app config and app token
	// perform compensating action if either fails
	return nil
}

// GetSingle fetches all data belonging to the app data
func (svc appadmin) GetSingle(ctx context.Context, in *proto.GetRequest) (*common.AppInfo, errors.ErrApi) {
	app, err := get.Single(ctx, svc.repo, in.GetAppUuid())
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not get App Information",
			Err:    err,
		}
	}
	return app, nil
}

// GetMultiple fetches all Apps related to the user asking for the data. The list contains
// only a minimal view with app-name and app-uuid
func (svc appadmin) GetMultiple(ctx context.Context, in *proto.GetListRequest) ([]*common.AppMetaInfo, errors.ErrApi) {

	apps, err := get.Multiple(ctx, svc.repo, in)
	if err != nil {
		if err == get.ErrNotFound {
			return nil, errors.ErrAPI{
				Status: http.StatusNotFound,
				Msg:    "Could not find any related Apps",
				Err:    err,
			}
		}
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not get Apps",
			Err:    err,
		}
	}
	return apps, nil
}

// MayAcquireToken verifies that the user trying to create an app token is allowed to do so
func (svc appadmin) MayAcquireToken(ctx context.Context, in *proto.MayAcquireTokenRequest) errors.ErrApi {
	ok, err := token.MayAcquire(ctx, svc.repo, in)
	if err != nil {
		if err == token.ErrNotFound {
			return errors.ErrAPI{Status: http.StatusBadRequest, Msg: "Could not find request App information", Err: err}
		}
		if err == token.ErrNotAuthorized {
			return errors.ErrAPI{Status: http.StatusUnauthorized, Msg: "Missing permission to acquire token", Err: err}
		}
		return errors.ErrAPI{Status: http.StatusInternalServerError, Msg: "Could not check for authorization", Err: err}
	}
	if !ok {
		return errors.ErrAPI{Status: http.StatusUnauthorized, Msg: "Missing permission to acquire token", Err: err}
	}
	return nil
}

func (svc appadmin) InviteToApp(ctx context.Context, in *proto.InviteRequest) (string, string, errors.ErrApi) {
	err := invite.ToApp(ctx, svc.repo, in.GetUserUuid(), in.GetOwnerUuid(), in.GetAppUuid())
	if err != nil {
		if err == invite.ErrNoAppFound {
			return "", "", errors.ErrAPI{
				Status: http.StatusBadRequest,
				Msg:    "Could not find requested app",
				Err:    err,
			}
		}
		return "", "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not invite user to app",
			Err:    err,
		}
	}
	app, err := get.Single(ctx, svc.repo, in.GetAppUuid())
	if err != nil {
		return "", "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not get App Information",
			Err:    err,
		}
	}
	return app.Name, app.Owner, nil
}
