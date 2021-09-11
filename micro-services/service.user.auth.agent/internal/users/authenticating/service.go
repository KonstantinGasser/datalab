package authenticating

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/initializing"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	Register(ctx context.Context, username, organization, password string) (string, errors.Api)
	Login(ctx context.Context, username, password string) (string, errors.Api)
	Authenticate(ctx context.Context, accessToken string) (*users.AuthedUser, errors.Api)
}

type service struct {
	repo           users.UserRepo
	permissionInit initializing.Service
}

func NewService(repo users.UserRepo, permissionInit initializing.Service) Service {
	return &service{
		repo:           repo,
		permissionInit: permissionInit,
	}
}

func (s service) Register(ctx context.Context, username, organization, password string) (string, errors.Api) {
	if ok, err := s.repo.UsernameTaken(ctx, username); err != nil || ok {
		if err != nil {
			return "", errors.New(http.StatusInternalServerError,
				err,
				"Could not create User Account")
		}
		if !ok {
			return "", errors.New(http.StatusBadRequest,
				fmt.Errorf("username taken"),
				"Username is already taken")
		}
	}

	newUser, err := users.NewDefaultUser(username, organization)
	if err != nil {
		if err == users.ErrInvalidOrgnName {
			return "", errors.New(http.StatusBadRequest,
				err,
				"Organization name includs invalid characters")
		}
		return "", errors.New(http.StatusBadRequest,
			err,
			"User Account requires a username, organization and password")
	}
	if err := newUser.HashAndSalt(password); err != nil {
		return "", errors.New(http.StatusInternalServerError,
			err,
			"Could not create User Account")
	}
	err = s.repo.Store(ctx, *newUser)
	if err != nil {
		return "", errors.New(http.StatusInternalServerError,
			err,
			"Could not create User Account")
	}

	// create init permissions for user
	if err := s.permissionInit.InitPermissions(ctx, newUser.Uuid, newUser.Organization); err != nil {
		return "", err
	}

	return newUser.Uuid, nil
}

func (s service) Login(ctx context.Context, username, password string) (string, errors.Api) {
	var storedUser users.User
	if err := s.repo.GetByUsername(ctx, username, &storedUser); err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New(http.StatusBadRequest,
				err,
				"Could not find any User with this Username")
		}
		return "", errors.New(http.StatusInternalServerError,
			err,
			"Could not get User data")
	}
	loginErr := storedUser.Credentials(password)
	if loginErr != nil {
		return "", errors.New(http.StatusUnauthorized,
			loginErr,
			"User credentials are wrong")
	}

	accessToken, err := storedUser.AccessToken()
	if err != nil {
		return "", errors.New(http.StatusInternalServerError,
			err,
			"Could not login User")
	}
	return accessToken, nil
}

func (s service) Authenticate(ctx context.Context, accessToken string) (*users.AuthedUser, errors.Api) {
	authedUser, err := users.LoggedIn(accessToken)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized,
			err,
			"Access Token is invalid")
	}
	return authedUser, nil
}
