package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) AuthUser(ctx context.Context, request *userSrv.AuthRequest) (*userSrv.AuthResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())

	logrus.Infof("<%v>[userService.AuthUser] received authentication request\n", ctx_value.GetString(ctx, "tracingID"))
	// status can be statusOK, statusInternalServerError or statusForbidden
	status, user, err := srv.user.Authenticate(ctx, srv.mongoClient, request.GetUsername(), request.GetPassword())
	if err != nil {
		logrus.Errorf("<%v>[userService.AuthUser] could not authenticate user:%v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.AuthResponse{
			StatusCode:    int32(status),
			Msg:           err.Error(),
			User:          nil,
			Authenticated: false,
		}, nil
	}
	logrus.Infof("<%v>[userService.AuthUser] user authenticated\n", ctx_value.GetString(ctx, "tracingID"))
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
