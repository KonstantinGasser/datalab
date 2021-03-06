package app

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func (app app) GetApps(ctx context.Context, mongo storage.Storage, req *appSrv.GetAppsRequest) ([]*appSrv.LightApp, error) {
	result, err := mongo.FindAll(ctx, dbName, appCollection, bson.D{
		{"ownerUUID", req.GetUserUuid()},
	})
	if err != nil {
		return nil, err
	}
	// parse bson.M slice to slice of lightApp as defined by the GetAppsResponse in appService grpc
	var appList []*appSrv.LightApp
	for _, res := range result {
		appList = append(appList, &appSrv.LightApp{
			Name: res["appName"].(string),
			Uuid: res["_id"].(string),
		})
	}
	return appList, nil
}
