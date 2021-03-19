package app

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func (app app) GetByID(ctx context.Context, mongo storage.Storage, req *appSrv.GetByIDRequest) (bson.M, error) {

	return mongo.FindOne(ctx, dbName, appCollection, bson.M{"_id": req.GetAppUuid()})
}
