package api

import (
	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/app"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
)

type AppService struct {
	appSrv.UnimplementedAppServiceServer
	mongoC storage.Storage
	app    app.App
}

func NewAppServiceServer() AppService {
	mongoC := storage.New("mongodb://AppDB:secure@192.168.178.163:27018")
	app := app.New()
	return AppService{
		mongoC: mongoC,
		app:    app,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv AppService) mustEmbedUnimplementedAppServiceServer() {}
