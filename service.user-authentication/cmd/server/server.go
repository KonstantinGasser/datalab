package server

import (
	"context"
	"net"

	"github.com/KonstantinGasser/datalab/service.user-authentication/domain"
	"github.com/KonstantinGasser/datalab/service.user-authentication/handler"
	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/service.user-authentication/repo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, host, dbAddr string) error {

	srv := grpc.NewServer()
	// create tokenService implementing the grpc.TokenServiceServer methods
	repo, err := repo.NewMongoDB(dbAddr)
	if err != nil {
		return err
	}
	domain := domain.NewUserAuthLogic(repo)

	userauthSvc := handler.NewHandler(domain)
	proto.RegisterUserAuthenticationServer(srv, userauthSvc)
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
