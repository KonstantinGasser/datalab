package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/sirupsen/logrus"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) AuthUser(ctx context.Context, request *userSrv.AuthRequest) (*userSrv.AuthResponse, error) {
	logrus.Info("[userService.AuthUser] received authentication request\n")
	// status can be statusOK, statusInternalServerError or statusForbidden
	status, user, err := srv.user.Authenticate(ctx, srv.mongoClient, request.GetUsername(), request.GetPassword())
	if err != nil {
		return &userSrv.AuthResponse{
			StatusCode:    int32(status),
			Msg:           err.Error(),
			User:          nil,
			Authenticated: false,
		}, nil
	}
	return &userSrv.AuthResponse{
		StatusCode: int32(status),
		Msg:        "user authenticated",
		User: &userSrv.AuthenticatedUser{
			Username:   user["username"].(string),
			Uuid:       user["_id"].(string),
			OrgnDomain: user["orgnDomain"].(string),
		},
		Authenticated: true,
	}, nil
}
