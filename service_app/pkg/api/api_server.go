package api

import (
	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	tokenSrv "github.com/KonstantinGasser/datalab/protobuf/token_service"
	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/service_app/pkg/app"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
)

type AppService struct {
	appSrv.UnimplementedAppServer
	storage storage.Storage
	app     app.App
	// *** Service Dependencies ***
	userService   userSrv.UserClient
	configService configSrv.ConfigClient
	tokenService  tokenSrv.TokenClient
}

func NewAppServer(store storage.Storage, user userSrv.UserClient, configService configSrv.ConfigClient, token tokenSrv.TokenClient) AppService {
	return AppService{
		storage:       store,
		app:           app.NewApp(),
		userService:   user,
		configService: configService,
		tokenService:  token,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv AppService) mustEmbedUnimplementedAppServiceServer() {}
