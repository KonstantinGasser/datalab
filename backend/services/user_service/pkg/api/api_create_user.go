package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/sirupsen/logrus"
)

// CreateUser receives the grpc request and handles user registration
func (srv UserService) CreateUser(ctx context.Context, request *userSrv.CreateUserRequest) (*userSrv.CreateUserResponse, error) {
	logrus.Info("[userService.CreateUser] received  create-user request\n")
	ctx, cancel := context.WithTimeout(context.Background(), queryTimeout)
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
