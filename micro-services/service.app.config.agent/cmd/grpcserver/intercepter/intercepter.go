package intercepter

import (
	"context"
	"fmt"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// TODO: must go in some ENV
type CtxKey string

const (
	secretUser           = "super_secure"
	secretService        = "super_secure"
	atuhCtxKey    CtxKey = "user"
)

var (
	ErrNotAuthenticated = fmt.Errorf("caller is not authenticated")
	ErrJWTParse         = fmt.Errorf("could not parse jwt token")
	ErrCorruptJWT       = fmt.Errorf("jwt could not be parsed (JWT might be corrupted)")
	ErrExpiredJWT       = fmt.Errorf("provided JWT has expired")
)

func WithUnary(interceptor, next grpc.UnaryServerInterceptor) grpc.ServerOption {
	// return grpc.UnaryInterceptor(middleware)
	return grpc.UnaryInterceptor(func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		return interceptor(ctx, req, info,
			func(nextCtx context.Context, nextReq interface{}) (interface{}, error) {
				return next(nextCtx, nextReq, info, handler)
			})
	})
}

func WithUserJwtAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod == "/config_proto.AppConfiguration/GetForClient" {
		return handler(ctx, req)
	}
	logrus.Info("[intercepter.WithUserJwtAuth] receiced request\n")

	authedUser, err := validateJWT(ctx)
	if err != nil {
		return nil, err
	}
	authCtx := context.WithValue(ctx, "user", authedUser)
	resp, err := handler(authCtx, req)

	return resp, err
}

func WithSvcJwtAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod != "/config_proto.AppConfiguration/GetForClient" {
		return handler(ctx, req)
	}
	logrus.Info("[intercepter.WithSvcJwtAuth] receiced request\n")

	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("could not find correct context details")
	}
	serviceToken, ok := meta["datalab-service-token"]
	if !ok {
		return nil, ErrNotAuthenticated
	}

	if _, err := verifyToken(serviceToken[0], secretService); err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func validateJWT(ctx context.Context) (*common.AuthedUser, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("could not find correct context details")
	}
	accessToken, ok := meta["authorization"]
	if !ok {
		return nil, ErrNotAuthenticated
	}
	token, err := verifyToken(accessToken[0], secretUser)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, ErrNotAuthenticated
	}
	if err := claims.Valid(); err != nil {
		return nil, ErrExpiredJWT
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
