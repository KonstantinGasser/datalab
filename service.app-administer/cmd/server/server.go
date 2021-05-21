package server

import (
	"context"
	"net"

	// appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	"github.com/KonstantinGasser/datalab/service.app-administer/domain"
	"github.com/KonstantinGasser/datalab/service.app-administer/grpc_adapter"
	"github.com/KonstantinGasser/datalab/service.app-administer/handler"
	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-administer/repo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, host, userAddr, configAddr, authAddr, apptokenAddr, dbAddr string) error {
	srv := grpc.NewServer()
	// create app-service
	// create app dependencies
	var err error
	repo, err := repo.NewMongoDB(dbAddr)
	userClient, err := grpc_adapter.NewUserAdministerClient(userAddr)
	configClient, err := grpc_adapter.NewAppConfigClient(configAddr)
	userauthClient, err := grpc_adapter.NewUserAuthClient(authAddr)
	apptokenClient, err := grpc_adapter.NewAppTokenIssuerClient(apptokenAddr)
	if err != nil {
		return err
	}
	appLogic := domain.NewAppLogic(repo, userClient, configClient, userauthClient, apptokenClient)
	appadminSvc := handler.NewHandler(appLogic)
	proto.RegisterAppAdministerServer(srv, appadminSvc)

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
