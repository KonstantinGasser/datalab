package api

import (
	"context"
	"net/http"

	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) Authenticate(ctx context.Context, in *userSrv.AuthenticateRequest) (*userSrv.AuthenticateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[userService.Authenticate] received request\n", ctx_value.GetString(ctx, "tracingID"))

	authedUser, err := srv.user.Authenticate(ctx, srv.storage, in.GetUsername(), in.GetPassword())
	if err != nil || authedUser == nil {
		logrus.Errorf("<%v>[userService.Authenticate] could not authenticate user or authedUser is <nil>: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &userSrv.AuthenticateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			UserClaims: nil,
		}, nil
	}
	return &userSrv.AuthenticateResponse{
		StatusCode: http.StatusOK,
		Msg:        "user is authenticated",
		UserClaims: &userSrv.UserTokenClaim{
			Uuid:       authedUser.UUID,
			OrgnDomain: authedUser.OrgnDomain,
		},
	}, nil
}
