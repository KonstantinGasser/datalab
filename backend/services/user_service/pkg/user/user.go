package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/repository"
	"github.com/KonstantinGasser/clickstream/utils/hash"
	"github.com/KonstantinGasser/clickstream/utils/unique"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	// dbUser is the name of user database in mongoDB
	dbUser = "datalabs_user"
	// collUser is the name of the collection used in dbUser
	// to store the user document
	collUser = "user"
)

type User struct{}

// Insert handles inserts of newly registered user into the mongo db.
// it checks if the username is already taken else calls the mongoClient.InsertOne to
// persist the user. It hashes the users password and assigns the user a UUID used as pk (_id)
// in MongoDB
func (user User) Insert(ctx context.Context, db *repository.MongoClient, username, password, orgnD string) (int, error) {
	// sanity check for user struct
	if username == "" || password == "" || orgnD == "" {
		return http.StatusBadRequest, fmt.Errorf("user information are missing")
	}
	// check if user already exists
	// errors of type mongo.ErrNoDocuments are excluded since they mean that no match
	// was found. an error here means the query failed
	resultMap, err := db.FindOne(ctx, dbUser, collUser, bson.M{"username": username})
	if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("could not execute FindOne: %v", err)
	}
	// mongos findOne query can return an empty bson.M struct if not found
	if len(resultMap) != 0 {
		return http.StatusBadRequest, fmt.Errorf("username already exists in system")
	}
	// primary-key (_id) for mongoDB document of user
	uuid, err := unique.UUID()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	hashedPassword, err := hash.FromPassword([]byte(password))
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
	// forward user byte slice to be persisted in DB/collection
	if err := db.InsertOne(ctx, dbUser, collUser, b); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// Authenticate checks whether passed credentials match records in the database in order to
// authenticate a user at login
func (user User) Authenticate(ctx context.Context, db *repository.MongoClient, username, password string) (int, bson.M, error) {

	result, err := db.FindOne(ctx, dbUser, collUser, bson.M{"username": username})
	if err != nil {
		return http.StatusInternalServerError, bson.M{}, fmt.Errorf("could not execute findOne: %v", err)
	}
	// if user is not found in the database (mongo.FindOne returns an empty bson.M struct)
	if len(result) == 0 {
		return http.StatusForbidden, bson.M{}, errors.New("could not find user in database")
	}
	if !hash.CheckPasswordHash(password, result["password"].(string)) {
		return http.StatusForbidden, bson.M{}, errors.New("user not authenticated")
	}
	// user is authenticated: returns user bson.M data
	return http.StatusOK, result, nil
}
