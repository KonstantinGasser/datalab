package user

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_user/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAll looks up all the records that match the provided UUIDs
func (user user) GetAll(ctx context.Context, storage storage.Storage, UUIDs []string) ([]UserItem, errors.ErrApi) {

	var userList []UserItem
	if err := storage.FindMany(ctx, userDatabase, userCollection, bson.M{"_id": bson.M{"$in": UUIDs}}, &userList); err != nil {
		// return http.StatusInternalServerError, nil, err
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not get User-List",
		}
	}
	return userList, nil
}
