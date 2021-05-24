package domain

import (
	"context"

	"github.com/KonstantinGasser/datalab/common"
	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	cfgsvc "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	apptokensvc "github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	usersvc "github.com/KonstantinGasser/datalab/service.user-administer/proto"
	userauthsvc "github.com/KonstantinGasser/datalab/service.user-authentication/proto"
)

type GatewayLogic interface {
	RegisterUser(ctx context.Context, form RegisterForm) errors.ErrApi
	LoginUser(ctx context.Context, form LoginForm) (string, errors.ErrApi)
	IsLoggedIn(ctx context.Context, token string) (*common.UserTokenClaims, errors.ErrApi)

	GetUserProfile(ctx context.Context, uuid string) (*common.UserInfo, errors.ErrApi)
	GetColleagues(ctx context.Context, userUuid string) ([]*common.UserInfo, errors.ErrApi)
	UpdateUserProfile(ctx context.Context, uuid string, form UserUpdateForm) errors.ErrApi

	CreateApp(ctx context.Context, uuid, organization string, form CreateAppForm) (string, errors.ErrApi)
	GetAppInfo(ctx context.Context, userUuid, appUuid string) (*common.AppInfo, errors.ErrApi)
	GetAppList(ctx context.Context, uuid string) ([]*common.AppMetaInfo, errors.ErrApi)

	CreateAppToken(ctx context.Context, userUuid, appUuid, appOrigin, appHash string) (*common.AppTokenInfo, errors.ErrApi)
	GetAppToken(ctx context.Context, uuid string) (*common.AppTokenInfo, errors.ErrApi)

	UpdateAppConfig(ctx context.Context, form UpdateConfigForm, flag string) errors.ErrApi
	GetAppConfig(ctx context.Context, uuid string) (*common.AppConfigInfo, errors.ErrApi)

	InviteToAppProcess(ctx context.Context, form InviteForm) errors.ErrApi
	AcceptInvite(ctx context.Context, form AcceptInviteForm) (string, errors.ErrApi)
}

type gatewaylogic struct {
	appClient       appsvc.AppAdministerClient
	apptokenClient  apptokensvc.AppTokenIssuerClient
	appconfigClient cfgsvc.AppConfigurationClient
	userClient      usersvc.UserAdministerClient
	userauthClient  userauthsvc.UserAuthenticationClient
}

func NewGatewayLogic(appC appsvc.AppAdministerClient, apptokenC apptokensvc.AppTokenIssuerClient,
	appcfgC cfgsvc.AppConfigurationClient, userC usersvc.UserAdministerClient,
	userauthC userauthsvc.UserAuthenticationClient) GatewayLogic {
	return &gatewaylogic{
		appClient:       appC,
		apptokenClient:  apptokenC,
		appconfigClient: appcfgC,
		userClient:      userC,
		userauthClient:  userauthC,
	}
}

type AppMetaData struct {
	Uuid, Origin, Hash string
}
