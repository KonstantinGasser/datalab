package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

<<<<<<< HEAD
func (srv UserService) GetUser(ctx context.Context, in *userSrv.GetUserRequest) (*userSrv.GetUserResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
=======
func (srv UserService) Get(ctx context.Context, request *userSrv.GetRequest) (*userSrv.GetResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())

>>>>>>> feature_app_token
	logrus.Infof("<%v>[userService.GetUser] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, user, err := srv.user.GetByID(ctx, srv.storage, in.GetForUuid())
	if err != nil {
		logrus.Errorf("<%v>[userService.GetUser] could not get user details: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.GetResponse{StatusCode: int32(status), Msg: err.Error(), User: nil}, nil
	}
	if status != 200 {
		return &userSrv.GetResponse{StatusCode: int32(status), Msg: err.Error(), User: nil}, nil
	}
	return &userSrv.GetResponse{
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
