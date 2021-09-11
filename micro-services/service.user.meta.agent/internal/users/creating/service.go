package creating

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users"
)

type Service interface {
	CreateUser(ctx context.Context, uuid, username, firstname, lastname, organization, position string) errors.Api
}

type service struct {
	repo users.UserRepository
}

func NewService(repo users.UserRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) CreateUser(ctx context.Context, uuid, username, firstname, lastname, organization, position string) errors.Api {
	newUser, err := users.NewDefault(uuid, username, firstname, lastname, organization, position)
	if err != nil {
		return errors.New(http.StatusBadRequest,
			err,
			"Mandatory Fields missing")
	}

	err = s.repo.Store(ctx, *newUser)
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not create new User")
	}
	return nil
}
