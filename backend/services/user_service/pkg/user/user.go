package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/repository"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct{}

// Insert inserts a new user into the mongo database
func (user User) Insert(ctx context.Context, db *repository.MongoClient, reqUser DBUser) (int, error) {
	// sanity check for user struct
	if reqUser.Username == "" || reqUser.Password == "" || reqUser.OrgnDomain == "" {
		return http.StatusBadRequest, fmt.Errorf("user information are missing")
	}
	// check if user already exists
	resultMap, err := db.FindOne(ctx, "datalabs_user", "user", bson.M{"username": reqUser.Username})
	if err != nil {
		logrus.Errorf("[user.Insert] could not execute FindOne: %v\n", err)
		return http.StatusInternalServerError, fmt.Errorf("could not execute FindOne: %v", err)
	}
	// if not 0 then user with username in db
	if len(resultMap) != 0 {
		return http.StatusBadRequest, fmt.Errorf("username already exists in system")
	}

	b, err := bson.Marshal(reqUser)
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("could not marshal MongoUser struct: %v", err)
	}
	if err := db.InsertOne(ctx, "datalabs_user", "user", b); err != nil {
		logrus.Errorf("[user.Insert] %s", err.Error())
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
