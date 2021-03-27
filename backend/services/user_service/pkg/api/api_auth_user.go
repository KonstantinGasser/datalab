package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) Authenticate(ctx context.Context, request *userSrv.AuthenticateRequest) (*userSrv.AuthenticateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[userService.AuthUser] received authentication request\n", ctx_value.GetString(ctx, "tracingID"))

	status, authedUser, err := srv.user.Authenticate(ctx, srv.storage, request.GetUsername(), request.GetPassword())
	if err != nil || authedUser == nil {
		logrus.Errorf("<%v>[userService.AuthUser] could not authenticate user or authedUser is <nil>: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.AuthenticateResponse{
			StatusCode: int32(status),
			Msg:        "could not authenticate user",
			UserClaims: nil,
		}, nil
	}
	return &userSrv.AuthenticateResponse{
		StatusCode: int32(status),
		Msg:        "user is authenticated",
		UserClaims: &userSrv.UserTokenClaim{
			Uuid:       authedUser.UUID,
			OrgnDomain: authedUser.OrgnDomain,
		},
	}, nil
}
