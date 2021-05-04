package server

import (
	"context"
	"net"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	"github.com/KonstantinGasser/datalab/service_app/pkg/api"
	"github.com/KonstantinGasser/datalab/service_app/pkg/grpcC"
	"github.com/KonstantinGasser/datalab/service_app/pkg/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, host, userAddr, tokenAddr, configAddr, dbAddr string) error {
	srv := grpc.NewServer()
	// create app-service
	// create app dependencies
	storage, err := storage.New(dbAddr)
	if err != nil {
		logrus.Fatalf("srv.Dependency] could not establish connection to storage:\n\t%v", err)
	}
	userClient, err := grpcC.NewUserClient(userAddr)
	if err != nil {
		logrus.Fatalf("srv.Dependency] could not establish connection to user-service:\n\t%v", err)
	}
	configClient, err := grpcC.NewConfigClient(configAddr)
	if err != nil {
		logrus.Fatalf("srv.Dependency] could not establish connection to config-service:\n\t%v", err)
	}
	tokenClient, err := grpcC.NewTokenClient(tokenAddr)
	if err != nil {
		logrus.Fatalf("srv.Dependency] could not establish connection to token-service:\n\t%v", err)
	}
	appService := api.NewAppServer(storage, userClient, configClient, tokenClient)
	appSrv.RegisterAppServer(srv, appService)

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
