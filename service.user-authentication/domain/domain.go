package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-configuration/errors"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/login"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/register"
	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/service.user-authentication/repo"
)

type UserAuthLogic interface {
	RegisterNewUser(ctx context.Context, in *proto.RegisterRequest) errors.ErrApi
	LoginUser(ctx context.Context, in *proto.LoginRequest) (string, errors.ErrApi)
	IsAuthenticated(ctx context.Context, in *proto.IsAuthedRequest) (*proto.Claims, errors.ErrApi)
}

type userauthlogic struct {
	repo repo.Repo
}

func NewUserAuthLogic(repo repo.Repo) UserAuthLogic {
	return &userauthlogic{
		repo: repo,
	}
}

func (svc userauthlogic) RegisterNewUser(ctx context.Context, in *proto.RegisterRequest) errors.ErrApi {
	err := register.NewUser(ctx, svc.repo, in)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create User-Account",
			Err:    err,
		}
	}
	return nil
}

func (svc userauthlogic) LoginUser(ctx context.Context, in *proto.LoginRequest) (string, errors.ErrApi) {
	token, err := login.User(ctx, svc.repo, in)
	if err != nil {
		if err == login.ErrUserNotFound {
			return "", errors.ErrAPI{
				Status: http.StatusNotFound,
				Msg:    "Could not find any User",
				Err:    err,
			}
		}
		if err == login.ErrWrongPassword {
			return "", errors.ErrAPI{
				Status: http.StatusUnauthorized,
				Msg:    "Username/Password are wrong",
				Err:    err,
			}
		}
	}
	return token, nil
}

func (svc userauthlogic) IsAuthenticated(ctx context.Context, in *proto.IsAuthedRequest) (*proto.Claims, errors.ErrApi) {
	claims, err := login.IsLoggedIn(ctx, in.GetJwt())
	if err != nil {
		if err == login.ErrCorruptedToken || err == login.ErrInvalidToken {
			return nil, errors.ErrAPI{
				Status: http.StatusUnauthorized,
				Msg:    "User is not authenticated",
				Err:    err,
			}
		}
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not verify if user is authenticated",
			Err:    err,
		}
	}
	return claims, nil
}
