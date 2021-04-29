package api

import (
	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	tokenSrv "github.com/KonstantinGasser/datalab/protobuf/token_service"
	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/service_app/pkg/app"
	"github.com/KonstantinGasser/datalab/service_app/pkg/grpcC"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
)

type AppService struct {
	appSrv.UnimplementedAppServer
	storage storage.Storage
	app     app.App
	// *** Service Dependencies ***
	userService  userSrv.UserClient
	tokenService tokenSrv.TokenClient
}

func NewAppServer(storage storage.Storage) AppService {
	app := app.NewApp()
	userService := grpcC.NewUserClient(":8001")
	tokenService := grpcC.NewTokenClient(":8002")
	return AppService{
		storage:      storage,
		app:          app,
		userService:  userService,
		tokenService: tokenService,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv AppService) mustEmbedUnimplementedAppServiceServer() {}
