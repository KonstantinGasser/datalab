package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_app/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_user/pkg/storage"
	"github.com/KonstantinGasser/datalab/utils/hash"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserItemAuth is a trimmed down version of the user since for authentication
// not all of the user data must be loaded from the database
type UserItemAuth struct {
	UUID       string `bson:"_id"`
	OrgnDomain string `bson:"orgn_domain"`
	Password   string `bson:"password"`
}

// Authenticate verifies that a user performing a login exists in the database and has provided the correct
// login credentials (username, password)
func (user user) Authenticate(ctx context.Context, storage storage.Storage, username, password string) (*UserItemAuth, errors.ErrApi) {

	var authedUser UserItemAuth
	if err := storage.FindOne(ctx, userDatabase, userCollection, bson.M{"username": username}, &authedUser); err != nil {
		if err == mongo.ErrNoDocuments { // as specified by the mongoClient.FindOne func
			return nil, errors.ErrAPI{
				Status: http.StatusBadRequest,
				Err:    fmt.Errorf("could not find any user in database"),
				Msg:    "Username does not exist",
			}
		}
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Err:    err,
			Msg:    "Could not authenticate user",
		}
	}
	if !hash.CheckPasswordHash(password, authedUser.Password) {
		return nil, errors.ErrAPI{
			Status: http.StatusUnauthorized,
			Err:    fmt.Errorf("password does not match db record"),
			Msg:    "Password is wrong",
		}
	}
	return &authedUser, nil
}
