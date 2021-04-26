package api

import (
	appSrv "github.com/KonstantinGasser/datalabs/protobuf/app_service"
	tokenSrv "github.com/KonstantinGasser/datalabs/protobuf/token_service"
	userSrv "github.com/KonstantinGasser/datalabs/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/service_app/pkg/app"
	"github.com/KonstantinGasser/datalabs/service_app/pkg/config"
	"github.com/KonstantinGasser/datalabs/service_app/pkg/grpcC"
	"github.com/KonstantinGasser/datalabs/service_app/pkg/storage"
)

type AppService struct {
	appSrv.UnimplementedAppServer
	storage storage.Storage
	app     app.App
	config  config.Config
	// *** Service Dependencies ***
	userService  userSrv.UserClient
	tokenService tokenSrv.TokenClient
}

func NewAppServer(storage storage.Storage) AppService {
	app := app.NewApp()
	config := config.NewConfig()
	userService := grpcC.NewUserClient(":8001")
	tokenService := grpcC.NewTokenClient(":8002")
	return AppService{
		storage:      storage,
		app:          app,
		config:       config,
		userService:  userService,
		tokenService: tokenService,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv AppService) mustEmbedUnimplementedAppServiceServer() {}
