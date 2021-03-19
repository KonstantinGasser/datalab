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
	mongoC storage.Storage
	app    app.App
	// *** Service Dependencies ***
	userService userSrv.UserServiceClient
}

func NewAppServiceServer() AppService {
	mongoC := storage.New("mongodb://AppDB:secure@192.168.0.179:27018")
	app := app.New()
	userService := grpcC.NewUserServiceClient(":8001")
	return AppService{
		mongoC:      mongoC,
		app:         app,
		userService: userService,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv AppService) mustEmbedUnimplementedAppServiceServer() {}
