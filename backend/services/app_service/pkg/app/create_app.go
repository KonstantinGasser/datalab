package app

import (
	"context"
	"errors"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateApp some docs
func (app app) CreateApp(ctx context.Context, mongo storage.Storage, appItem AppItem) (int, error) {

	// duplicate names may exists in the system but owners can only hold unique app names
	exists, err := mongo.Exists(ctx, appDatabase, appCollection, bson.M{"appName": appItem.AppName, "ownerUUID": appItem.OwnerUUID})
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if exists {
		return http.StatusBadRequest, errors.New("duplicated app names are not possible")
	}

	if err := mongo.InsertOne(ctx, appDatabase, appCollection, appItem); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
