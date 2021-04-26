package app

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// HasPermissions verifies that a given caller of a request is indeed allowed to retrieve or modify the passed app data
func (app app) HasPermissions(ctx context.Context, storage storage.Storage, callerUUID, appUUID string) (int, bool, error) {
	filter := bson.D{
		{
			Key: "$and",
			Value: bson.A{
				bson.M{"_id": appUUID},
				bson.M{"owner_uuid": callerUUID},
			},
		},
	}
	ok, err := storage.Exists(ctx, appDatabase, appCollection, filter)
	if err != nil {
		return http.StatusInternalServerError, false, err
	}
	return http.StatusOK, ok, nil
}
