package collecting

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client"
	"github.com/KonstantinGasser/required"
	"github.com/sirupsen/logrus"
)

type Service interface {
	GetApp(ctx context.Context, r *apps.GetAppRequest) *apps.GetAppResponse
	GetAppList(ctx context.Context, r *apps.GetAppListRequest) *apps.GetAppListResponse
}

type service struct {
	appMetaClient   client.ClientAppMeta
	userMetaClient  client.ClientUserMeta
	appTokenClient  client.ClientAppToken
	appConfigClient client.ClientAppConfig
}

func NewService(
	appMetaClient client.ClientAppMeta,
	userMetaClient client.ClientUserMeta,
	appTokenClient client.ClientAppToken,
	appConfigClient client.ClientAppConfig,
) Service {
	return &service{
		appMetaClient:   appMetaClient,
		userMetaClient:  userMetaClient,
		appTokenClient:  appTokenClient,
		appConfigClient: appConfigClient,
	}
}

func (s service) GetApp(ctx context.Context, r *apps.GetAppRequest) *apps.GetAppResponse {
	if err := required.Atomic(r); err != nil {
		return &apps.GetAppResponse{
			Status: http.StatusOK,
			Msg:    "Mandatory fields missing",
			Err:    err.Error(),
		}
	}

	app, err := s.appMetaClient.GetApp(ctx, r)
	if err != nil {
		return &apps.GetAppResponse{
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}
	token, conf, owner, collectErr := s.collectAttachedAppData(ctx, app.Uuid, r.AuthedUser)
	if collectErr != nil {
		logrus.Errorf("[app.collection.Get] could not get all data: %v\n", collectErr)
	}
	return &apps.GetAppResponse{
		Status: http.StatusOK,
		Msg:    "App Data",
		App:    app,
		Config: conf,
		Token:  token,
		Owner:  owner,
	}

}

func (s service) GetAppList(ctx context.Context, r *apps.GetAppListRequest) *apps.GetAppListResponse {
	applist, err := s.appMetaClient.GetAppList(ctx, r)
	if err != nil {
		return &apps.GetAppListResponse{
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}
	return &apps.GetAppListResponse{
		Status: http.StatusOK,
		Msg:    "User's Apps",
		Apps:   applist,
	}
}

func (s service) collectAttachedAppData(ctx context.Context, appUuid string, authedUser *common.AuthedUser) (*common.AppAccessToken, *common.AppConfigurations, *common.UserInfo, error) {
	withCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	var errC = make(chan error)
	var resC = make(chan struct {
		Field string
		Value interface{}
	})
	go s.appTokenClient.CollectAppToken(withCancel, appUuid, authedUser, resC, errC)
	go s.appConfigClient.CollectAppConfig(withCancel, appUuid, authedUser, resC, errC)
	go s.userMetaClient.CollectOwnerInfo(withCancel, authedUser, resC, errC)

	var apptoken *common.AppAccessToken
	var appconfig *common.AppConfigurations
	var appowner *common.UserInfo
	for i := 0; i < 3; i++ {
		select {
		case err := <-errC:
			if err != nil {
				logrus.Errorf("[%s][creating.EmitInit] emit cause error: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
				// if there is an error while emitting events
				// here the emiited events must succed in order for the
				// transaction to succeed - hence if err cancel context and
				// role back (if that would have been implmeneted)
				// return nil, nil, nil, err
				continue
			}
		case result := <-resC:
			switch result.Field {
			case "apptoken":
				// fmt.Println("token: ", result)
				var ok bool
				apptoken, ok = result.Value.(*common.AppAccessToken)
				if !ok { // if assertion fails value will be nil which is not nice but sometimes will happen thou
					continue
				}
			case "appconfig":
				var ok bool
				appconfig, ok = result.Value.(*common.AppConfigurations)
				if !ok { // if assertion fails value will be nil which is not nice but sometimes will happen thou
					continue
				}
			case "appowner":
				// fmt.Println("owner: ", result)
				var ok bool
				appowner, ok = result.Value.(*common.UserInfo)
				if !ok { // if assertion fails value will be nil which is not nice but sometimes will happen thou
					continue
				}
			}
		}
	}
	close(errC)
	close(resC)
	return apptoken, appconfig, appowner, nil
}
