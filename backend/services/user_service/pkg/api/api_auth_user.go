package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) Authenticate(ctx context.Context, in *userSrv.AuthenticateRequest) (*userSrv.AuthenticateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[userService.Authenticate] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, authedUser, err := srv.user.Authenticate(ctx, srv.storage, in.GetUsername(), in.GetPassword())
	if err != nil || authedUser == nil {
		logrus.Errorf("<%v>[userService.Authenticate] could not authenticate user or authedUser is <nil>: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
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
