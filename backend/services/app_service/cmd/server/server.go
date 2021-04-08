package server

import (
	"context"
	"net"

	appSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/app_service"
	"github.com/KonstantinGasser/datalabs/backend/services/app_service/pkg/api"
	"github.com/KonstantinGasser/datalabs/backend/services/app_service/pkg/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, addr string) error {
	srv := grpc.NewServer()
	// create app-service
	appService := api.NewAppServer(storage.New("mongodb://AppDB:secure@192.168.0.179:27018"))
	appSrv.RegisterAppServer(srv, appService)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("[server.Run] cloud not listen to %s: %v", addr, err)
	}
	logrus.Infof("[server.Run] listening on %s\n", addr)

	err = srv.Serve(listener)
	// if context gets canceled in main (invoked by some SIG)
	// perform graceful shutdown of server
	defer srv.GracefulStop()
	if err != nil {
		return err
	}
	return nil
}
