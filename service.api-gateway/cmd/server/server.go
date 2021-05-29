package server

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api-gateway/adapter"
	"github.com/KonstantinGasser/datalab/service.api-gateway/domain"
	"github.com/KonstantinGasser/datalab/service.api-gateway/handler"
	"github.com/sirupsen/logrus"
)

// Run acts as a run abstraction to start the HTTP-Server
// and can return an error - which is nice when called
// from main
func Run(ctx context.Context, hostAddr, userAddr, appAddr, apptokenAddr, tokenAddr, configAddr string) error {
	// create api dependencies
	userClient, err := adapter.CreateUserAdminClient(userAddr)
	if err != nil {
		return err
	}
	appClient, err := adapter.CreateAppAdminClient(appAddr)
	if err != nil {
		return err
	}
	apptokenClient, err := adapter.CreateAppTokenIssuerClient(apptokenAddr)
	if err != nil {
		return err
	}
	appconfigClient, err := adapter.CreateAppConfigClient(configAddr)
	if err != nil {
		return err
	}
	userauthClient, err := adapter.CreateUserAuthClient(tokenAddr)
	if err != nil {
		return err
	}

	domain := domain.NewGatewayLogic(appClient, apptokenClient, appconfigClient, userClient, userauthClient)

	gatewaysvc := handler.NewHandler(domain)
	logrus.Info("[api.Dependency] established connection to all services\n")
	gatewaysvc.Apply(handler.WithAllowedOrigins("*"))

	gatewaysvc.Register("/api/v1/user/register", gatewaysvc.RegisterUser, /// Done
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
	)
	gatewaysvc.Register("/api/v1/user/login", gatewaysvc.LoginUser, /// Done
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
	)

	gatewaysvc.Register("/api/v1/user/profile", gatewaysvc.GetUserProfile, /// Done
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
		gatewaysvc.WithAuth,
	)

	gatewaysvc.Register("/api/v1/user/profile/colleagues", gatewaysvc.GetUserColleagues, /// Done
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
		gatewaysvc.WithAuth,
	)

	gatewaysvc.Register("/api/v1/user/profile/update", gatewaysvc.UpdateUserProfile, /// Done
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
		gatewaysvc.WithAuth,
	)

	gatewaysvc.Register("/api/v1/app/create", gatewaysvc.CreateApp, /// Done
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
		gatewaysvc.WithAuth,
	)
	gatewaysvc.Register("/api/v1/app/get", gatewaysvc.GetAppDetails, /// Done
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
		gatewaysvc.WithAuth,
	)
	gatewaysvc.Register("/api/v1/app/getall", gatewaysvc.GetAppList, /// Done
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
		gatewaysvc.WithAuth,
	)

	gatewaysvc.Register("/api/v1/app/token/create", gatewaysvc.CreateAppToken, /// Done
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
		gatewaysvc.WithAuth,
	)

	gatewaysvc.Register("/api/v1/app/config/upsert", gatewaysvc.UpdateAppConfig, /// Done
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
		gatewaysvc.WithAuth,
	)
	gatewaysvc.Register("/api/v1/app/member/invite", gatewaysvc.InviteUserToApp,
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
		gatewaysvc.WithAuth,
	)
	gatewaysvc.Register("/api/v1/app/member/invite/accept", gatewaysvc.AcceptInvite,
		gatewaysvc.WithTracing,
		gatewaysvc.WithCors,
		gatewaysvc.WithAuth,
	)

	// waiting for context to be canceled
	// not implemented: graceful shutdown
	go func() {
		<-ctx.Done()
		logrus.Infof("Server cleaning up...")
	}()
	logrus.Infof("[server.Run] listening on %s\n", hostAddr)
	if err := http.ListenAndServe(hostAddr, nil); err != nil {
		return err
	}
	return nil
}
