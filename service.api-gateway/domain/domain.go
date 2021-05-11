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
	IsLoggedIn(ctx context.Context, token string) (*common.TokenClaims, errors.ErrApi)

	GetUserProfile(ctx context.Context, uuid string) (*common.UserInfo, errors.ErrApi)
	UpdateUserProfile(ctx context.Context, uuid string, form UserUpdateForm) errors.ErrApi
}

type gatewaylogic struct {
	appClient       appsvc.AppAdministerClient
	apptokenClient  apptokensvc.AppTokenIssuerClient
	appconfigClient cfgsvc.AppConfigurationClient
	userClient      usersvc.UserAdministerClient
	userauthClient  userauthsvc.UserAuthenticationClient
}

func NewGatewayLogic(appC appsvc.AppAdministerClient, apptokenC apptokensvc.AppTokenIssuerClient, appcfgC cfgsvc.AppConfigurationClient, userC usersvc.UserAdministerClient, userauthC userauthsvc.UserAuthenticationClient) GatewayLogic {
	return &gatewaylogic{
		appClient:       appC,
		apptokenClient:  apptokenC,
		appconfigClient: appcfgC,
		userClient:      userC,
		userauthClient:  userauthC,
	}
}
