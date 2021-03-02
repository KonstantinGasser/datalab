package api

import (
	"context"
	"fmt"
	"net/http"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
)

func (srv UserService) AuthUser(ctx context.Context, request *userSrv.AuthRequest) (*userSrv.AuthResponse, error) {

	user, err := srv.user.Authenticate(ctx, srv.mongoClient, request.GetUsername(), request.GetPassword())
	if err != nil {
		return &userSrv.AuthResponse{
			StatusCode:    http.StatusForbidden,
			Msg:           "user not authenticated",
			User:          nil,
			Authenticated: false,
		}, fmt.Errorf("[userService.AuthUser] could not authenticate user: %v", err)
	}
	return &userSrv.AuthResponse{
		StatusCode: http.StatusOK,
		Msg:        "user authenticated",
		User: &userSrv.AuthenticatedUser{
			Username:   user["username"].(string),
			Uuid:       user["_id"].(string),
			OrgnDomain: user["orgnDomain"].(string),
		},
		Authenticated: true,
	}, nil
}
