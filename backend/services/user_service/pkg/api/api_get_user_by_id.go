package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv UserService) GetUserByID(ctx context.Context, request *userSrv.GetUserByIDRequest) (*userSrv.GetUserByIDResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())

	logrus.Infof("<%v>[userService.GetUserByID] received get user by id request\n", ctx_value.GetString(ctx, "tracingID"))

	status, user, err := srv.user.GetByID(ctx, srv.storage, request.GetUuid())
	if err != nil {
		logrus.Errorf("<%v>[userService.GetUserByID] could not get user details: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.GetUserByIDResponse{StatusCode: int32(status), Msg: err.Error(), User: nil}, nil
	}
	if status != 200 {
		return &userSrv.GetUserByIDResponse{StatusCode: int32(status), Msg: err.Error(), User: nil}, nil
	}
	logrus.Warn(user.ProfileImgURL)
	return &userSrv.GetUserByIDResponse{
		StatusCode: int32(status),
		Msg:        "requested user found",
		User: &userSrv.User{
			Uuid:          user.UUID,
			Username:      user.Username,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			OrgnDomain:    user.OrgnDomain,
			OrgnPosition:  user.OrgnPosition,
			ProfileImgUrl: user.ProfileImgURL,
		},
	}, nil
}
