package server

import (
	"context"
	"net"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/api"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/repository"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

// Run is a run-abstraction for the main func
func Run(ctx context.Context, addr string) error {
	srv := grpc.NewServer()

	// create new UserService with all dependencies
	mongoC, err := repository.NewMongoClient("mongodb://userDB:secure@192.168.0.179:27017")
	errorFatal(err)
	userService := api.NewUserService(mongoC)

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

// errorFatal is used to reduce the if err != nil statements and will
// fail fatally if error is not nil
// only used for dependencies
func errorFatal(err error) {
	if err == nil {
		return
	}
	logrus.Fatalf("%v", err)
}
