package server

import (
	"context"
	"net"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/api"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/apptoken"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/grpcC"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, host, appAddr, dbAddr string) error {
	srv := grpc.NewServer()
	// create storage dependency
	apptoken := apptoken.New()
	storage, err := storage.New(dbAddr)
	if err != nil {
		return err
	}
	appClient, err := grpcC.NewAppClient(appAddr)
	if err != nil {
		return err
	}
	apptokenService := api.NewAppTokenServer(apptoken, appClient, storage)
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
