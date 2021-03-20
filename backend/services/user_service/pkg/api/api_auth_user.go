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

	status, authedUser, err := srv.user.Authenticate(ctx, srv.storage, request.GetUsername(), request.GetPassword())
	if err != nil || authedUser == nil {
		logrus.Errorf("<%v>[userService.AuthUser] could not authenticate user or authedUser is <nil>: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.AuthResponse{
			StatusCode:    int32(status),
			Msg:           "could not authenticate user",
			Authenticated: false,
			User:          nil,
		}, nil
	}
	return &userSrv.AuthResponse{
		StatusCode:    int32(status),
		Msg:           "user is authenticated",
		Authenticated: true,
		User: &userSrv.AuthenticatedUser{
			Uuid:     authedUser.UUID,
			Username: request.GetUsername(),
			// these fields will be dropped in the next version since the user object must only
			// hold data relevant for the JWT not the other data!
			OrgnDomain:    "",
			FirstName:     "",
			LastName:      "",
			OrgnPosition:  "",
			ProfileImgUrl: "",
		},
	}, nil
}
