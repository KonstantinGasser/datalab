package domain

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/create"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/delete"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/get"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain/token"
	"github.com/KonstantinGasser/datalab/service.app-administer/errors"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	cfgsvc "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	aptissuer "github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	usersvc "github.com/KonstantinGasser/datalab/service.user-administer/proto"
)

type AppAdmin interface {
	Create(ctx context.Context, in *proto.CreateRequest) (string, errors.ErrApi)
	Delete(ctx context.Context, in *proto.DeleteRequest) errors.ErrApi
	GetSingle(ctx context.Context, in *proto.GetRequest) (*common.AppInfo, *common.AppConfigInfo, *common.AppTokenInfo, errors.ErrApi)
	GetMultiple(ctx context.Context, in *proto.GetListRequest) ([]*common.AppMetaInfo, errors.ErrApi)
	MayAcquireToken(ctx context.Context, in *proto.MayAcquireTokenRequest) errors.ErrApi
}

type appadmin struct {
	userSvc     usersvc.UserAdministerClient
	configSvc   cfgsvc.AppConfigurationClient
	tkissuerSvc aptissuer.AppTokenIssuerClient
	repo        repo.Repo
}

func NewAppLogic(repo repo.Repo, user usersvc.UserAdministerClient, config cfgsvc.AppConfigurationClient, token aptissuer.AppTokenIssuerClient) AppAdmin {
	return &appadmin{}
}

// Create handles the creation of a new app coordinating the initialization of the app-config
// and performs a rollback if initialization fails
func (svc appadmin) Create(ctx context.Context, in *proto.CreateRequest) (string, errors.ErrApi) {
	appUuid, err := create.App(ctx, svc.repo, in)
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create new App",
			Err:    err,
		}
	}
	if resp, err := svc.configSvc.Init(ctx, &cfgsvc.InitRequest{
		Tracing_ID: in.GetTracing_ID(),
		ForApp:     appUuid,
	}); err != nil || resp.GetStatusCode() != http.StatusOK {
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
	return appUuid, nil
}

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

func (svc appadmin) GetSingle(ctx context.Context, in *proto.GetRequest) (*common.AppInfo, *common.AppConfigInfo, *common.AppTokenInfo, errors.ErrApi) {

	var wg sync.WaitGroup
	wg.Add(1)
	var cfgResp *cfgsvc.GetResponse
	var cfgErr error
	go func() {
		cfgResp, cfgErr = svc.configSvc.Get(ctx, &cfgsvc.GetRequest{Tracing_ID: in.GetTracing_ID(), ForUuid: in.GetAppUuid()})
	}()

	wg.Add(1)
	var aptResp *aptissuer.GetResponse
	var aptErr error
	go func() {
		aptResp, aptErr = svc.tkissuerSvc.Get(ctx, &aptissuer.GetRequest{Tracing_ID: in.GetTracing_ID(), AppUuid: in.GetAppUuid()})
	}()
	if cfgErr != nil || aptErr != nil {
		return nil, nil, nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not get App Information",
			Err:    fmt.Errorf("%v | %v", cfgErr, aptErr),
		}
	}
	app, err := get.Single(ctx, svc.repo, in)
	if err != nil {
		return nil, nil, nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not get App Information",
			Err:    err,
		}
	}
	return app, cfgResp.GetConfigs(), aptResp.GetToken(), nil
}

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

func (svc appadmin) MayAcquireToken(ctx context.Context, in *proto.MayAcquireTokenRequest) errors.ErrApi {
	ok, err := token.MayAcquire(ctx, svc.repo, in)
	if err != nil {
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
