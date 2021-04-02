package user

import (
	"context"

	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/storage"
)

// User interface describes the APIs which can be used to interact with
// the UserService
type User interface {
	Register(ctx context.Context, storage storage.Storage, in RegisterRequest) (int, error)
	Login(ctx context.Context, storage storage.Storage, in LoginRequest) (int, error)
}

func New() User {
	return &user{}
}

// Defined Request struct allowing for type safety

// RegisterRequest represents the passed HTTP JSON by the client
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginRequest represents the passed HTTP JSON by the client
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
