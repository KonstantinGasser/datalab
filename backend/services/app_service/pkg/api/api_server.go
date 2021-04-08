package api

import (
	appSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/app_service"
	tokenSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/token_service"
	userSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/backend/services/app_service/pkg/app"
	"github.com/KonstantinGasser/datalabs/backend/services/app_service/pkg/grpcC"
	"github.com/KonstantinGasser/datalabs/backend/services/app_service/pkg/storage"
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
