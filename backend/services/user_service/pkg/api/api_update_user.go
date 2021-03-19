package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) UpdateUser(ctx context.Context, request *userSrv.UpdateUserRequest) (*userSrv.UpdateUserResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[userService.UpdateUser] received update user request\n", ctx_value.GetString(ctx, "tracingID"))

	status, err := srv.user.Update(ctx, srv.mongoClient, request)
	if err != nil {
		logrus.Errorf("<%v>[userService.UpdateUser] could not update user: %v\n", err)
		return &userSrv.UpdateUserResponse{
			StatusCode: int32(status),
			Msg:        err.Error(),
		}, nil
	}
	return &userSrv.UpdateUserResponse{
		StatusCode: int32(status),
		Msg:        "user account updated",
	}, nil

}