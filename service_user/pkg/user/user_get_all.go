package user

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_user/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAll looks up all the records that match the provided UUIDs
func (user user) GetAll(ctx context.Context, storage storage.Storage, UUIDs []string) (int, []UserItem, error) {

	var userList []UserItem
	if err := storage.FindMany(ctx, userDatabase, userCollection, bson.M{"_id": bson.M{"$in": UUIDs}}, &userList); err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, userList, nil
}
