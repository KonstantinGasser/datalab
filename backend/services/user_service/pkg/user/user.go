package user

import (
	"context"

	"github.com/KonstantinGasser/datalabs/backend/services/user_service/pkg/storage"
)

const (
	// userDatabase is the name of the mongoDB
	userDatabase = "datalabs_user"
	// userCollection is the name of the collection used to
	// store user documents
	userCollection = "user"
)

type User interface {
	// Create inserts a new user into the database
	Create(ctx context.Context, storage storage.Storage, userItem UserItem) (int, error)

	// Authenticate verifies whether a user's credentials match with the ones stored in the database
	Authenticate(ctx context.Context, storage storage.Storage, username, password string) (int, *UserItemAuth, error)

	// Update updates a userItem in the database
	Update(ctx context.Context, storage storage.Storage, userItem UserItemUpdateable) (int, error)

	// GetAll collects all user details for all given UUIDs
	GetAll(ctx context.Context, storage storage.Storage, UUIDs []string) (int, []UserItem, error)

	// Get collects all user details for the given UUID
	Get(ctx context.Context, storage storage.Storage, UUID string) (int, UserItem, error)

	// CompareOrgn compares users based on some indicator
	CompareOrgn(ctx context.Context, storage storage.Storage, baseObject string, compareWith []string) (int, bool, []string, error)
}

type user struct{}

func NewUser() User {
	return &user{}
}

// UserItem is a representation of a user document which lives in the
// mongoDB. Fields must be exported in order to serve as (de-)serialization for the mongoDB
type UserItem struct {
	UUID          string `bson:"_id" required:"yes"`
	Username      string `bson:"username" required:"yes"`
	Password      string `bson:"password" required:"yes"`
	FirstName     string `bson:"first_name" required:"yes"`
	LastName      string `bson:"last_name" required:"yes"`
	OrgnDomain    string `bson:"orgn_domain" required:"yes"`
	OrgnPosition  string `bson:"orgn_position" required:"yes"`
	ProfileImgURL string `bson:"profile_img_url" required:"yes"`
}

// UserItemUpdateable describes the fields of a user which can be updated by the user
type UserItemUpdateable struct {
	UUID          string `bson:"_id" required:"yes"`
	FirstName     string `bson:"first_name" required:"yes"`
	LastName      string `bson:"last_name" required:"yes"`
	OrgnPosition  string `bson:"orgn_position" required:"yes"`
	ProfileImgURL string `bson:"profile_img_url" required:"yes"`
}
