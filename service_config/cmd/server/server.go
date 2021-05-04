package server

import (
	"context"
	"net"

	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	"github.com/KonstantinGasser/datalab/service_config/pkg/api"
	"github.com/KonstantinGasser/datalab/service_config/pkg/config"
	"github.com/KonstantinGasser/datalab/service_config/pkg/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, host, dbAddr string) error {
	srv := grpc.NewServer()
	// create storage dependency
	storage, err := storage.New(dbAddr)
	if err != nil {
		return err
	}
	config := config.New()
	configService := api.NewConfigServer(config, storage)
	configSrv.RegisterConfigServer(srv, configService)

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
