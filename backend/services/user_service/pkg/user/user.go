package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/repository"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/utils"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct{}

// Insert inserts a new user into the mongo database
func (user User) Insert(ctx context.Context, db *repository.MongoClient, username, password, orgnD string) (int, error) {
	// sanity check for user struct
	if username == "" || password == "" || orgnD == "" {
		return http.StatusBadRequest, fmt.Errorf("user information are missing")
	}
	// check if user already exists
	resultMap, err := db.FindOne(ctx, "datalabs_user", "user", bson.M{"username": username})
	if err != nil {
		logrus.Errorf("[user.Insert] could not execute FindOne: %v\n", err)
		return http.StatusInternalServerError, fmt.Errorf("could not execute FindOne: %v", err)
	}
	// if not 0 then user with username in db
	if len(resultMap) != 0 {
		return http.StatusBadRequest, fmt.Errorf("username already exists in system")
	}

	uuid, err := utils.UUID()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	hashedPassword, err := utils.HashFromPassword([]byte(password))
	if err != nil {
		return http.StatusInternalServerError, err
	}
	b, err := bson.Marshal(newDBUser(
		uuid,
		username,
		hashedPassword,
		orgnD,
	))
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("could not marshal MongoUser struct: %v", err)
	}
	if err := db.InsertOne(ctx, "datalabs_user", "user", b); err != nil {
		logrus.Errorf("[user.Insert] %s", err.Error())
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (user User) Authenticate(ctx context.Context, db *repository.MongoClient, username, password string) (bson.M, error) {

	result, err := db.FindOne(ctx, "datalabs_user", "user", bson.M{"username": username})
	if err != nil || len(result) == 0 {
		return bson.M{}, fmt.Errorf("could not execute findOne: %v", err)
	}

	if utils.CheckPasswordHash(password, result["password"].(string)) {
		return bson.M{}, errors.New("user not authenticated")
	}
	// user is authenticated: returns user map
	return result, nil
}
