package app

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteApp deletes an app based on the provided appUUID
func (app app) DeleteApp(ctx context.Context, mongo storage.Storage, appUUID string) (int, error) {
	if err := mongo.DeleteOne(ctx, appDatabase, appCollection, bson.M{"_id": appUUID}); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
