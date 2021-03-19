package app

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func (app app) GetApps(ctx context.Context, mongo storage.Storage, req *appSrv.GetAppsRequest) ([]*appSrv.LightApp, error) {

	var appList []AppItemLight
	if err := mongo.FindAll(ctx, dbName, appCollection, bson.D{{"owner_uuid", req.GetUserUuid()}}, &appList); err != nil {
		return nil, err
	}
	// convert AppItemLight to appSrv.LightApp slice
	var resultList []*appSrv.LightApp = make([]*appSrv.LightApp, len(appList))
	for i, item := range appList {
		resultList[i] = &appSrv.LightApp{
			Name: item.AppName,
			Uuid: item.UUID,
		}
	}
	return resultList, nil
}
