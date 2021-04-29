package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_user/pkg/storage"
	"go.mongodb.org/mongo-driver/bson"
)

// Create inserts a new UserItem into the mongoDB. Checks before that the user name is not already taken if so returns
// an http.StatusBadeRequest, error
func (user user) Create(ctx context.Context, storage storage.Storage, userItem UserItem) errors.ErrApi {
	// check if provided organization names match the criteria
	ok := orgnAllowed(userItem.OrgnDomain)
	if !ok {
		return errors.ErrAPI{
			Status: http.StatusBadRequest,
			Err:    fmt.Errorf("organization name not allowed"),
			Msg:    "Organization must not have any of the following chars [/]",
		}
	}
	// check if username is taken
	taken, err := storage.Exists(ctx, userDatabase, userCollection, bson.M{"_id": userItem.UUID})

	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not create account",
		}
	}
	if taken { // username exists but must be unique
		return errors.ErrAPI{
			Status: http.StatusBadRequest,
			Err:    fmt.Errorf("username is already taken"),
			Msg:    "Username is already taken",
		}
	}
	// inserts new user in storage
	if err := storage.InsertOne(ctx, userDatabase, userCollection, userItem); err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not create account",
		}
	}
	return nil
}
