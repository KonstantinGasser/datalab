package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/KonstantinGasser/datalabs/service_user/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// Create inserts a new UserItem into the mongoDB. Checks before that the user name is not already taken if so returns
// an http.StatusBadeRequest, error
func (user user) Create(ctx context.Context, storage storage.Storage, userItem UserItem) (int, error) {
	// check if username is taken
	taken, err := storage.Exists(ctx, userDatabase, userCollection, bson.M{"_id": userItem.UUID})
	// check if provided organization names match the criteria
	ok := orgnAllowed(userItem.OrgnDomain)
	if !ok {
		return http.StatusBadRequest, errors.New("organization domain must not include either of (/,)")
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if taken { // username exists but must be unique
		return http.StatusBadGateway, errors.New("username already taken")
	}
	// inserts new user in storage
	if err := storage.InsertOne(ctx, userDatabase, userCollection, userItem); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
