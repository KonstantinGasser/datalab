package app

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// GetByID collects all the app details for a given appUUID
func (app app) GetByID(ctx context.Context, mongo storage.Storage, appUUID string) (int, AppItem, error) {

	var appDetails AppItem
	err := mongo.FindOne(ctx, appDatabase, appCollection, bson.M{"_id": appUUID}, &appDetails)
	if err != nil {
		return http.StatusInternalServerError, AppItem{}, err
	}
	return http.StatusOK, appDetails, nil
}
