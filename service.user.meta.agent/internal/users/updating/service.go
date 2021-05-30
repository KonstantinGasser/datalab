package updating

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users"
)

type Service interface {
	UpadeUser(ctx context.Context, uuid, firstname, lastname, position string) errors.Api
}

type service struct {
	repo users.UserRepository
}

func NewService(repo users.UserRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) UpadeUser(ctx context.Context, uuid, firstname, lastname, position string) errors.Api {

	updatable := users.UpdatableUser{
		Uuid:      uuid,
		FirstName: firstname,
		LastName:  lastname,
		Position:  position,
	}
	err := s.repo.UpdateUser(ctx, updatable)
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not create new User")
	}
	return nil
}
