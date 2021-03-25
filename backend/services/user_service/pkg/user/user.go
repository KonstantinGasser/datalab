package user

import (
	"context"

	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/storage"
)

const (
	// userDatabase is the name of the mongoDB
	userDatabase = "datalabs_user"
	// userCollection is the name of the collection used to
	// store user documents
	userCollection = "user"
)

type User interface {
	// InsertNew inserts a new user into the database
	InsertNew(ctx context.Context, storage storage.Storage, userItem UserItem) (int, error)
	// Authenticate verifies whether a user's credentials match with the ones stored in the database
	Authenticate(ctx context.Context, storage storage.Storage, username, password string) (int, *UserItemAuth, error)
	// Update updates a userItem in the database
	Update(ctx context.Context, storage storage.Storage, userItem UserItemUpdateable) (int, error)
	// GetByIDs collects all user details for all given UUIDs
	GetByIDs(ctx context.Context, storage storage.Storage, UUIDs []string) (int, []UserItem, error)
	// GetByID collects all user details for the given UUID
	GetByID(ctx context.Context, storage storage.Storage, UUID string) (int, UserItem, error)
	// CompareUsers compares users based on some indicator
	Compare(ctx context.Context, storage storage.Storage, comparator Comparator, comparable Comparable) (int, *CompareResult, error)
}

type user struct{}

func NewUser() User {
	return &user{}
}
