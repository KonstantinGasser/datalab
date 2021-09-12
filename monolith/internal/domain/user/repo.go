package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrNotFound        = fmt.Errorf("could not find user")
	ErrDuplicatedEntry = fmt.Errorf("user already exists")
)

type Repository interface {
	Store(ctx context.Context, user User) error
	Update(ctx context.Context, updatable UpdatableUser) error
	ById(ctx context.Context, uuid uuid.UUID, stored interface{}) error
	ByOrganization(ctx context.Context, organization string, stored interface{}) error
}
