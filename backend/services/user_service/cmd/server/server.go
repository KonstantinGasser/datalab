package server

import (
	"context"
	"net"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/api"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

// Run is a run-abstraction for the main func
func Run(ctx context.Context, addr string) error {
	srv := grpc.NewServer()

	// create new UserService with all dependencies
	userService, err := api.NewUserService()
	if err != nil {
		logrus.Fatalf("[server.Run] could not create api.UserService: %v", err)
	}
	// register grpc server to service
	userSrv.RegisterUserServiceServer(srv, userService)
	// create tcp listener
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("[server.Run] cloud not listen to %s: %v", addr, err)
	}
	logrus.Infof("[server.Run] listening on %s\n", addr)

	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}
