package app

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func (app app) DeleteApp(ctx context.Context, mongo storage.Storage, req *appSrv.DeleteAppRequest) (int, error) {

	if err := mongo.DeleteOne(ctx, dbName, appCollection, bson.D{
		{"_id", req.GetAppUuid()},
	}); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
