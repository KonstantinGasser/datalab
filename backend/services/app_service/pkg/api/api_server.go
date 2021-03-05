package api

import (
	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
)

type AppService struct {
	appSrv.UnimplementedAppServiceServer
	mongoC storage.Storage
}

func NewAppServiceServer() AppService {
	mongoC := storage.New("mongodb://sdfsdfsdf")
	return AppService{
		mongoC: mongoC,
	}
}

// this is required to implemented due to the new go-grpc update: may change in the future
func (srv AppService) mustEmbedUnimplementedAppServiceServer() {}
