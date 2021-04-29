package app

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// AddMember adds all provided []member to the app if the the requested caller is the owner of the app
// must ensure that requested members belong to the same organization as the app owner
func (app app) AddMember(ctx context.Context, storage storage.Storage, ownerUUID, appUUID string, member []string) errors.ErrApi {

	// filter must ensure that caller has permissions (aka is owner) of the app
	filterAppAndOwner := bson.M{
		"_id":        appUUID,
		"owner_uuid": ownerUUID,
	}

	updateQuery := bson.D{
		{
			Key: "$addToSet",
			Value: bson.M{
				"member": bson.M{
					"$each": member,
				},
			},
		},
	}
	// updated shows if documents have been updated or not
	updated, err := storage.UpdateOne(ctx, appDatabase, appCollection, filterAppAndOwner, updateQuery, false)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not add member(s) to app",
		}
	}

	if updated == 0 {
		// not yet sure what to do with this information
		// return http.StatusForbidden, errors.New("user not permitted to modify app data")
	}
	return nil
}
