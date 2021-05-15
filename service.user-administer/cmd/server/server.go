package server

import (
	"context"
	"net"

	"github.com/KonstantinGasser/datalab/service.user-administer/domain"
	"github.com/KonstantinGasser/datalab/service.user-administer/grpc_adapter"
	"github.com/KonstantinGasser/datalab/service.user-administer/handler"
	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/service.user-administer/repo"
	"github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

// Run is a run-abstraction for the main func
func Run(ctx context.Context, host, permissionsAddr, dbAddr string) error {
	srv := grpc.NewServer()
	// create storage dependency
	repo, err := repo.NewMongoDB(dbAddr)
	if err != nil {
		return err
	}
	permissionClient, err := grpc_adapter.NewUserPermissionsClient(permissionsAddr)
	if err != nil {
		return err
	}
	domain := domain.NewUserAdminLogic(repo, permissionClient)
	userAdminSvc := handler.NewHandler(domain)
	proto.RegisterUserAdministerServer(srv, userAdminSvc)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		logrus.Fatalf("[server.Run] cloud not listen to %s: %v", host, err)
	}
	logrus.Infof("[server.Run] listening on %s\n", host)

	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}
