package user

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalabs/service_user/pkg/storage"
	"github.com/KonstantinGasser/datalabs/utils/hash"
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
func (user user) Authenticate(ctx context.Context, storage storage.Storage, username, password string) (int, *UserItemAuth, error) {

	var authedUser UserItemAuth
	if err := storage.FindOne(ctx, userDatabase, userCollection, bson.M{"username": username}, &authedUser); err != nil {
		if err == mongo.ErrNoDocuments { // as specified by the mongoClient.FindOne func
			return http.StatusForbidden, nil, nil
		}
		return http.StatusInternalServerError, nil, err
	}
	if !hash.CheckPasswordHash(password, authedUser.Password) {
		return http.StatusForbidden, nil, nil
	}
	return http.StatusOK, &authedUser, nil
}
