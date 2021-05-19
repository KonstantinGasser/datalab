package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-configuration/errors"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/login"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/permissions"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/register"
	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/types"
	"github.com/KonstantinGasser/datalab/service.user-authentication/jwts"
	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/service.user-authentication/repo"
)

// UserAuthLogic is the interface for this service
type UserAuthLogic interface {
	RegisterNewUser(ctx context.Context, in *proto.RegisterRequest) (string, errors.ErrApi)
	LoginUser(ctx context.Context, in *proto.LoginRequest) (string, errors.ErrApi)
	IsAuthenticated(ctx context.Context, in *proto.IsAuthedRequest) (*common.UserTokenClaims, errors.ErrApi)
	AddAppAccess(ctx context.Context, in *proto.AddAppAccessRequest) (string, errors.ErrApi)
}

type userauthlogic struct {
	repo repo.Repo
}

func NewUserAuthLogic(repo repo.Repo) UserAuthLogic {
	return &userauthlogic{
		repo: repo,
	}
}

// RegisterNewUser coordinates the use-case of registering a new user
func (svc userauthlogic) RegisterNewUser(ctx context.Context, in *proto.RegisterRequest) (string, errors.ErrApi) {
	uuid, err := register.NewUser(ctx, svc.repo, in)
	if err != nil {
		if err == register.ErrUserAlreadyExists {
			return "", errors.ErrAPI{
				Status: http.StatusBadRequest,
				Msg:    "Username already exists",
				Err:    err,
			}
		}
		if err == register.ErrInvalidOrgnName {
			return "", errors.ErrAPI{
				Status: http.StatusBadRequest,
				Msg:    "Organization Name must not contain a forward-slash",
				Err:    err,
			}
		}
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create User-Account",
			Err:    err,
		}
	}
	return uuid, nil
}

// LoginUser coordinates the use-case of login in a user
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

// IsAuthenticated handles logic concerning the authentication of a user
func (svc userauthlogic) IsAuthenticated(ctx context.Context, in *proto.IsAuthedRequest) (*common.UserTokenClaims, errors.ErrApi) {
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

func (svc userauthlogic) AddAppAccess(ctx context.Context, in *proto.AddAppAccessRequest) (string, errors.ErrApi) {
	permission := types.AppPermission{
		AppUuid: in.GetAppUuid(),
		Role:    types.AppRole(in.GetAppRole()),
	}
	newPermission, err := permissions.UpdateAppAccess(ctx, svc.repo, in.GetUserUuid(), permission)
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not add App Permission",
			Err:    err,
		}
	}
	token, err := jwts.Issue(newPermission.UserUuid, newPermission.UserOrgn, newPermission.Apps)
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not add App Permission",
			Err:    err,
		}
	}
	return token, nil
}
