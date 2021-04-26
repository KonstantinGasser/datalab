package user

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_user/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// GetByID collects all user details for the given user uuid
func (user user) Get(ctx context.Context, storage storage.Storage, UUID string) (int, UserItem, error) {
	var userItem UserItem
	if err := storage.FindOne(ctx, userDatabase, userCollection, bson.M{"_id": UUID}, &userItem); err != nil {
		return http.StatusInternalServerError, UserItem{}, nil
	}
	return http.StatusOK, userItem, nil
}
