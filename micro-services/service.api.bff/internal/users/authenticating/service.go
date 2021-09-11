package authenticating

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client"
	"github.com/KonstantinGasser/required"
	"github.com/dgrijalva/jwt-go"
)

const (
	secretUser = "super_secure"
)

var (
	ErrNotAuthenticated = fmt.Errorf("user is not authenticated")
	ErrJWTParse         = fmt.Errorf("could not parse jwt token")
	ErrCorruptJWT       = fmt.Errorf("jwt could not be parsed (JWT might be corrupted)")
	ErrExpiredJWT       = fmt.Errorf("provided JWT has expired")
)

type Service interface {
	Register(ctx context.Context, r *users.RegisterRequest) *users.RegisterResponse
	Login(ctx context.Context, r *users.LoginRequest) *users.LoginResponse
	Authenticate(ctx context.Context, accessToken string) (*common.AuthedUser, errors.Api)
}

type service struct {
	userAuthClient client.ClientUserAuth
	userMetaClient client.ClientUserMeta
}

func NewService(userAuthClient client.ClientUserAuth, userMetaClient client.ClientUserMeta) Service {
	return &service{
		userAuthClient: userAuthClient,
		userMetaClient: userMetaClient,
	}
}

func (s service) Register(ctx context.Context, r *users.RegisterRequest) *users.RegisterResponse {
	if err := required.Atomic(r); err != nil {
		return &users.RegisterResponse{
			Status: http.StatusBadRequest,
			Msg:    "Missing mandatory fields",
			Err:    err.Error(),
		}
	}

	userUuid, err := s.userAuthClient.Register(ctx, r)
	if err != nil {
		return &users.RegisterResponse{
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}
	r.UserUuid = userUuid
	err = s.userMetaClient.CreateUserProfile(ctx, r)
	if err != nil {
		return &users.RegisterResponse{
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}
	return &users.RegisterResponse{
		Status: http.StatusOK,
		Msg:    "User Account created",
	}
}

func (s service) Login(ctx context.Context, r *users.LoginRequest) *users.LoginResponse {
	if err := required.Atomic(r); err != nil {
		return &users.LoginResponse{
			Status: http.StatusBadRequest,
			Msg:    "Username and Password required",
			Err:    err.Error(),
		}
	}

	accessToken, err := s.userAuthClient.Login(ctx, r)
	if err != nil {
		return &users.LoginResponse{
			Status: err.Code(),
			Msg:    err.Info(),
			Err:    err.Error(),
		}
	}
	return &users.LoginResponse{
		Status:      http.StatusOK,
		Msg:         "User logged in",
		AccessToken: accessToken,
	}
}

func (s service) Authenticate(ctx context.Context, accessToken string) (*common.AuthedUser, errors.Api) {

	token, err := verifyToken(accessToken, secretUser)
	if err != nil {
		return nil, errors.New(http.StatusUnauthorized, ErrNotAuthenticated, "User not authenticated")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New(http.StatusUnauthorized, ErrNotAuthenticated, "User not authenticated")
	}
	if err := claims.Valid(); err != nil {
		return nil, errors.New(http.StatusUnauthorized, ErrExpiredJWT, "User not authenticated")
	}

	return &common.AuthedUser{
		Uuid:         claims["sub"].(string),
		Organization: claims["orgn"].(string),
		Username:     claims["uname"].(string),
	}, nil
}

func verifyToken(tokenString, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCorruptJWT
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, ErrJWTParse
	}
	return token, nil
}
