package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/storage"
)

type user struct{}

// Register registers a new user in a given storage. The new user must have a unique username else operation
// will fail
func (srv user) Register(ctx context.Context, storage storage.Storage, in RegisterRequest) (int, error) {
	// check that username is unique
	if ok := storage.Exists(ctx, in.Username); ok {
		return http.StatusBadRequest, errors.New("username must be unique")
	}

	if err := storage.Put(ctx, in.Username, in.Password); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// Login verifies that the login credentials match with the database record
func (srv user) Login(ctx context.Context, storage storage.Storage, in LoginRequest) (int, error) {
	password, err := storage.Get(ctx, in.Username)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if password != in.Password {
		return http.StatusUnauthorized, nil
	}
	return http.StatusOK, nil
}
