package api

import (
	"context"
	"net/http"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv UserService) GetUsersByID(ctx context.Context, request *userSrv.GetUsersByIDRequest) (*userSrv.GetUsersByIDResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())

	logrus.Infof("<%v>[userService.GetUsersByID] received get users by id request\n", ctx_value.GetString(ctx, "tracingID"))

	users, err := srv.user.GetByIDs(ctx, srv.mongoClient, request.GetUserUuids())
	if err != nil {
		logrus.Errorf("<%v>[userService.GetUsersByID] could not execute GetByIDs: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.GetUsersByIDResponse{
			StatusCode: http.StatusInternalServerError,
			Msg:        "Could not get users information",
			Users:      []*userSrv.User{},
		}, nil
	}
	return &userSrv.GetUsersByIDResponse{
		StatusCode: http.StatusOK,
		Msg:        "users record by uuids",
		Users:      users,
	}, nil
}
