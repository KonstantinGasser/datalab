package fetching

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	FetchById(ctx context.Context, uuid string) (*users.User, errors.Api)
	FetchByOrganization(ctx context.Context, organization string) ([]users.User, errors.Api)
}

type service struct {
	repo users.UserRepository
}

func NewService(repo users.UserRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) FetchById(ctx context.Context, uuid string) (*users.User, errors.Api) {
	var user users.User
	if err := s.repo.GetById(ctx, uuid, &user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusNotFound,
				err,
				"Could not find User data")
		}
	}
	return &user, nil
}

func (s service) FetchByOrganization(ctx context.Context, organization string) ([]users.User, errors.Api) {
	var userList []users.User
	if err := s.repo.GetByOrganization(ctx, organization, &userList); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusNotFound,
				err,
				"Could not find Users data")
		}
	}
	return userList, nil
}
