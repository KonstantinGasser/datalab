package server

import (
	"context"
	"net"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/grpc_adapter"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/handler"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/repo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, host, appAddr, dbAddr string) error {
	srv := grpc.NewServer()
	// create storage dependency
	repo, err := repo.NewMongoDB(dbAddr)
	if err != nil {
		return err
	}
	appadminClient, err := grpc_adapter.NewAppAdministerClient(appAddr)
	if err != nil {
		return err
	}
	domain := domain.NewAppTokenLogic(repo, appadminClient)
	apptokenService := handler.NewHandler(domain)
	proto.RegisterAppTokenIssuerServer(srv, apptokenService)

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
