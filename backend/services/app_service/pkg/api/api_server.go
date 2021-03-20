package api

import (
	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/api_gateway/pkg/grpcC"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/app"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
)

type AppService struct {
	appSrv.UnimplementedAppServiceServer
	storage storage.Storage
	app     app.App
	// *** Service Dependencies ***
	userService userSrv.UserServiceClient
}

func NewAppServiceServer(storage storage.Storage) AppService {
	app := app.NewApp()
	userService := grpcC.NewUserServiceClient(":8001")
	return AppService{
		storage:     storage,
		app:         app,
		userService: userService,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv AppService) mustEmbedUnimplementedAppServiceServer() {}
