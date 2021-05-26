package server

import (
	"context"
	"net"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/dao"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/handler"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, host, dbAddr string) error {
	srv := grpc.NewServer()
	// create storage dependency
	repo, err := dao.NewMongoDB(dbAddr)
	if err != nil {
		return err
	}

	domain := domain.NewDomainLogic(repo)
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
