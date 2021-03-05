package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/utils"
	"github.com/sirupsen/logrus"
)

// CreateUser receives the grpc request and handles user registration
func (srv UserService) CreateUser(ctx context.Context, request *userSrv.CreateUserRequest) (*userSrv.CreateUserResponse, error) {
	// add tracingID from request to context
	ctx = utils.AddValCtx(ctx, "tracingID", request.GetTracing_ID())

	logrus.Info("<%v>[userService.CreateUser] received  create-user request\n", utils.StringValueCtx(ctx, "tracingID"))
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()
	status, err := srv.user.Insert(ctx, srv.mongoClient, request.GetUsername(), request.GetPassword(), request.GetOrgnDomain())
	if err != nil {
		return &userSrv.CreateUserResponse{
			StatusCode: int32(status),
			Msg:        err.Error(),
		}, nil
	}
	return &userSrv.CreateUserResponse{
		StatusCode: int32(status),
		Msg:        "user added to system",
	}, nil
}
