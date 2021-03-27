package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv UserService) GetUser(ctx context.Context, request *userSrv.GetUserRequest) (*userSrv.GetUserResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())

	logrus.Infof("<%v>[userService.GetUser] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, user, err := srv.user.GetByID(ctx, srv.storage, request.GetForUuid())
	if err != nil {
		logrus.Errorf("<%v>[userService.GetUser] could not get user details: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.GetUserResponse{StatusCode: int32(status), Msg: err.Error(), User: nil}, nil
	}
	if status != 200 {
		return &userSrv.GetUserResponse{StatusCode: int32(status), Msg: err.Error(), User: nil}, nil
	}
	return &userSrv.GetUserResponse{
		StatusCode: int32(status),
		Msg:        "requested user found",
		User: &userSrv.ComplexUser{
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
