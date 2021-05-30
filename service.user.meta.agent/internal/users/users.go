package users

import (
	"context"

	"github.com/KonstantinGasser/required"
)

type UserRepository interface {
	Store(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, updatable UpdatableUser) error
	GetById(ctx context.Context, uuid string, stored interface{}) error
	GetByOrganization(ctx context.Context, organization string, stored interface{}) error
}

type User struct {
	Uuid         string `bson:"_id"`
	Username     string `bson:"username" required:"yes" min:"5"`
	FirstName    string `bson:"first_name" required:"yes"`
	LastName     string `bson:"last_name" required:"yes"`
	Organization string `bson:"organization" required:"yes"`
	Position     string `bson:"position" required:"yes"`
}

// UpdatableUser defins the fields of a User that can be changed
type UpdatableUser struct {
	Uuid                          string
	FirstName, LastName, Position string
}

func NewDefault(uuid, username, firstname, lastname, organization, position string) (*User, error) {
	user := &User{
		Uuid:         uuid,
		Username:     username,
		FirstName:    firstname,
		LastName:     lastname,
		Organization: organization,
		Position:     position,
	}
	if err := required.Atomic(user); err != nil {
		return nil, err
	}
	return user, nil
}
