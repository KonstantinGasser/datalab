package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// HasPermissions verifies that a given caller of a request is indeed allowed to retrieve or modify the passed app data
func (app app) HasPermissions(ctx context.Context, storage storage.Storage, callerUUID, appUUID string) errors.ErrApi {
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
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not authorize request",
		}
	}
	if !ok {
		return errors.ErrAPI{
			Status: http.StatusUnauthorized,
			Err:    fmt.Errorf("caller has no permissions for this resource"),
			Msg:    "Missing permissions for action",
		}
	}
	return nil
}
