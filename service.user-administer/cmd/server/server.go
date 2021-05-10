package server

import (
	"context"
	"net"

	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/service_user/pkg/api"
	"github.com/KonstantinGasser/datalab/service_user/pkg/storage"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

// Run is a run-abstraction for the main func
func Run(ctx context.Context, host, dbAddr string) error {
	srv := grpc.NewServer()

	// create new UserService
	// database dependency to mongoDB
	mongoC := storage.NewMongoClient(dbAddr)
	userService := api.NewUserService(mongoC)

	// register grpc server to service
	userSrv.RegisterUserServer(srv, userService)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		logrus.Fatalf("[server.Run] cloud not listen to %s: %v", host, err)
	}
	logrus.Infof("[server.Run] listening on %s\n", host)

	err = srv.Serve(listener)
	// if context gets canceled in main (invoked by some SIG)
	// perform graceful shutdown of server
	defer srv.GracefulStop()
	if err != nil {
		return err
	}
	return nil
}

// errorFatal is used to reduce the if err != nil statements and will
// fail fatally if error is not nil
// *** only used for dependencies at boot-up ***
func errorFatal(err error) {
	if err == nil {
		return
	}
	logrus.Fatalf("%v", err)
}
