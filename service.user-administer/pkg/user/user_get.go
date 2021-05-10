package user

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_user/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// GetByID collects all user details for the given user uuid
func (user user) Get(ctx context.Context, storage storage.Storage, UUID string) (UserItem, errors.ErrApi) {
	var userItem UserItem
	if err := storage.FindOne(ctx, userDatabase, userCollection, bson.M{"_id": UUID}, &userItem); err != nil {
		return UserItem{}, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not get User-Details",
		}
	}
	return userItem, nil
}
