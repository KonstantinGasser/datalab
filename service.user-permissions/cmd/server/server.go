package server

import (
	"context"
	"net"

	"github.com/KonstantinGasser/datalab/service.user-permissions/domain"
	"github.com/KonstantinGasser/datalab/service.user-permissions/handler"
	"github.com/KonstantinGasser/datalab/service.user-permissions/proto"
	"github.com/KonstantinGasser/datalab/service.user-permissions/repo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, host, dbAddr string) error {
	srv := grpc.NewServer()
	// create storage dependency
	repo, err := repo.NewMongoDB(dbAddr)
	if err != nil {
		return err
	}
	domain := domain.NewPermissionsLogic(repo)
	handler := handler.NewHandler(domain)
	proto.RegisterUserPermissionsServer(srv, handler)

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
