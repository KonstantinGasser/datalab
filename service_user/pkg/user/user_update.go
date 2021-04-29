package user

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_user/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// Update updates the allowed fields of the user record
func (user user) Update(ctx context.Context, storage storage.Storage, userItem UserItemUpdateable) errors.ErrApi {

	updateQuery := bson.D{
		{
			Key:   "$set",
			Value: userItem,
		},
	}
	if err := storage.UpdateOne(ctx, userDatabase, userCollection, bson.M{"_id": userItem.UUID}, updateQuery); err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not update account",
		}
	}
	return nil
}
