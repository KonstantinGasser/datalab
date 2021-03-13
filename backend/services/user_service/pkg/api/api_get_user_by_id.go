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

	status, user, err := srv.user.GetByID(ctx, srv.mongoClient, request.GetUuid())
	if err != nil {
		return &userSrv.GetUserByIDResponse{
			StatusCode: int32(status),
			Msg:        err.Error(),
			User:       nil,
		}, nil
	}
	if len(user) == 0 {
		return &userSrv.GetUserByIDResponse{
			StatusCode: int32(status),
			Msg:        err.Error(),
			User:       nil,
		}, nil
	}
	return &userSrv.GetUserByIDResponse{
		StatusCode: int32(status),
		Msg:        "requested user found",
		User: &userSrv.AuthenticatedUser{
			Username:     user["username"].(string),
			FirstName:    user["first_name"].(string),
			LastName:     user["last_name"].(string),
			OrgnDomain:   user["orgnDomain"].(string), // change in db to orgn_domain mongo js style
			OrgnPosition: user["orgn_position"].(string),
		},
	}, nil
}
