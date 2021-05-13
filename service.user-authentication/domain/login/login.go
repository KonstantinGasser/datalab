package login

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/service.user-authentication/config"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/login/jwts"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/service.user-authentication/repo"
	"github.com/KonstantinGasser/datalab/utils/hash"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrUserNotFound  = fmt.Errorf("could not find any user")
	ErrWrongPassword = fmt.Errorf("provided password does not match records")
)

// User checks if the provided user credentials match with the database records
func User(ctx context.Context, repo repo.Repo, in *proto.LoginRequest) (string, error) {

	var foundUser types.UserAuthInfo
	err := repo.FindOne(ctx, config.UserAuthDB, config.UserAuthColl, bson.M{"username": in.GetUsername()}, &foundUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", ErrUserNotFound
		}
		return "", err
	}
	if !hash.CheckPasswordHash(in.GetPassword(), foundUser.Password) {
		return "", ErrWrongPassword
	}

	token, err := jwts.Issue(foundUser.Uuid, foundUser.Organization)
	if err != nil {
		return "", err
	}
	return token, nil
}
