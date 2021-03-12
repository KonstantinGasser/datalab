package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// CreateUser receives the grpc request and handles user registration
func (srv UserService) CreateUser(ctx context.Context, request *userSrv.CreateUserRequest) (*userSrv.CreateUserResponse, error) {
	// add tracingID from request to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())

	logrus.Infof("<%v>[userService.CreateUser] received  create-user request\n", ctx_value.GetString(ctx, "tracingID"))
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()
	logrus.Info(request)
	status, err := srv.user.Insert(ctx, srv.mongoClient,
		request.GetUsername(),
		request.GetPassword(),
		request.GetOrgnDomain(),
		request.GetFirstName(),
		request.GetLastName(),
		request.GetOrgnPosition(),
	)
	if err != nil {
		logrus.Errorf("<%v>[userService.CreateUser] could not create user:%v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.CreateUserResponse{
			StatusCode: int32(status),
			Msg:        err.Error(),
		}, nil
	}
	logrus.Infof("<%v>[userService.CreateUser] user created\n", ctx_value.GetString(ctx, "tracingID"))
	return &userSrv.CreateUserResponse{
		StatusCode: int32(status),
		Msg:        "user added to system",
	}, nil
}
