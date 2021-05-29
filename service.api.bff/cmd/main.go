package main

import (
	"flag"
	"log"

	"github.com/KonstantinGasser/datalab/service.api.bff/cmd/httpserver"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps/collecting"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps/creating"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps/inviting"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps/modifying"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users/authenticating"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users/fetching"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users/updating"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client"
	"github.com/sirupsen/logrus"
)

func main() {
	host := flag.String("host", "localhost:8080", "address to run the server on")
	userMetaAddr := flag.String("user-srv", "localhost:8001", "address to connect to user-service")
	appMetaAddr := flag.String("app-srv", "localhost:8003", "address to connect to app-service")
	apptokenAddr := flag.String("apptoken-srv", "localhost:8006", "address to connect to app-service")
	appconfigAddr := flag.String("config-srv", "localhost:8005", "address to connect to app-service")
	userauthAddr := flag.String("token-srv", "localhost:8002", "address to connect to token-service")
	flag.Parse()

	// create grpc service dependencies
	grpcUserAuth, err := client.NewClientUserAuth(*userauthAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	grpcUserMeta, err := client.NewClientUserMeta(*userMetaAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	grpcAppMeta, err := client.NewClientAppMeta(*appMetaAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	grpcAppToken, err := client.NewClientAppToken(*apptokenAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	grpcAppConfig, err := client.NewClientAppConfig(*appconfigAddr)
	if err != nil {
		logrus.Fatal(err)
	}

	// create server dependencies
	userauthSerivce := authenticating.NewService(*grpcUserAuth, *grpcUserMeta)
	userupdateService := updating.NewService(*grpcUserMeta)
	userfetchService := fetching.NewService(*grpcUserMeta)

	appcreateService := creating.NewService(*grpcAppMeta, *grpcAppToken)
	appcollectService := collecting.NewService(*grpcAppMeta, *grpcUserMeta, *grpcAppToken, *grpcAppConfig)
	appmodifyService := modifying.NewService(*grpcAppConfig)
	appinviteService := inviting.NewService(*grpcAppMeta)

	server := httpserver.NewDefault(
		userauthSerivce,
		userupdateService,
		userfetchService,
		appcreateService,
		appcollectService,
		appmodifyService,
		appinviteService,
	)

	server.Apply(httpserver.WithAllowedOrigins("*"))

	// setting up routes
	server.Register("/api/v1/user/register", server.RegisterUser,
		server.WithTracing,
		server.WithCors,
	)
	server.Register("/api/v1/user/login", server.LoginUser,
		server.WithTracing,
		server.WithCors,
	)
	server.Register("/api/v1/user/profile/update", server.UpdateUserProfile,
		server.WithTracing,
		server.WithCors,
		server.WithAuth,
	)
	server.Register("/api/v1/user/profile", server.GetUserProfile,
		server.WithTracing,
		server.WithCors,
		server.WithAuth,
	)
	server.Register("/api/v1/user/colleagues", server.GetColleagues,
		server.WithTracing,
		server.WithCors,
		server.WithAuth,
	)

	server.Register("/api/v1/app/create", server.CreateApp,
		server.WithTracing,
		server.WithCors,
		server.WithAuth,
	)
	server.Register("/api/v1/app", server.GetApp,
		server.WithTracing,
		server.WithCors,
		server.WithAuth,
	)
	server.Register("/api/v1/app/all", server.GetAppList,
		server.WithTracing,
		server.WithCors,
		server.WithAuth,
	)
	server.Register("/api/v1/app/token/issue", server.IssueAppToken,
		server.WithTracing,
		server.WithCors,
		server.WithAuth,
	)

	server.Register("/api/v1/app/config/update", server.UpdateAppConfig,
		server.WithTracing,
		server.WithCors,
		server.WithAuth,
	)
	server.Register("/api/v1/app/invite", server.SendAppInvite,
		server.WithTracing,
		server.WithCors,
		server.WithAuth,
	)
	server.Register("/api/v1/app/invite/accept", server.AcceptInvite,
		server.WithTracing,
		server.WithCors,
		server.WithAuth,
	)
	// run server
	log.Fatal(server.Start(*host))
}
