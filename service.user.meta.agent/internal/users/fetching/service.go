package fetching

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	FetchLoggedIn(ctx context.Context) (*users.User, errors.Api)
	FetchById(ctx context.Context, uuid string) (*users.User, errors.Api)
	FetchByOrganization(ctx context.Context) ([]users.User, errors.Api)
}

type service struct {
	repo users.UserRepository
}

func NewService(repo users.UserRepository) Service {
	return &service{
		repo: repo,
	}
}

func (s service) FetchLoggedIn(ctx context.Context) (*users.User, errors.Api) {
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return nil, errors.New(http.StatusUnauthorized, fmt.Errorf("missing authentication"), "User not authenticated")
	}

	var user users.User
	if err := s.repo.GetById(ctx, authedUser.Uuid, &user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusNotFound,
				err,
				"Could not find User data")
		}
	}
	return &user, nil
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

func (s service) FetchByOrganization(ctx context.Context) ([]users.User, errors.Api) {
	authedUser, ok := ctx.Value("user").(*common.AuthedUser)
	if !ok {
		return nil, errors.New(http.StatusUnauthorized, fmt.Errorf("missing authentication"), "User not authenticated")
	}
	var userList []users.User
	if err := s.repo.GetByOrganization(ctx, authedUser.Organization, &userList); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusNotFound,
				err,
				"Could not find Users data")
		}
	}
	return userList, nil
}
