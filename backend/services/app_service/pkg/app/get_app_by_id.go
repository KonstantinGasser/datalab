package app

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func (app app) GetByID(ctx context.Context, mongo storage.Storage, req *appSrv.GetByIDRequest) (AppItem, error) {

	var appDetails AppItem
	logrus.Info(req.GetAppUuid())
	err := mongo.FindOne(ctx, dbName, appCollection, bson.M{"_id": req.GetAppUuid()}, &appDetails)
	if err != nil {
		return AppItem{}, err
	}
	return appDetails, nil
}
