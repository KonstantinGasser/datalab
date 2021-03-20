package server

import (
	"context"
	"net"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/api"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, addr string) error {
	srv := grpc.NewServer()
	// create app-service
	appService := api.NewAppServiceServer(storage.New("mongodb://AppDB:secure@192.168.0.179:27018"))
	appSrv.RegisterAppServiceServer(srv, appService)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatalf("[server.Run] cloud not listen to %s: %v", addr, err)
	}
	logrus.Infof("[server.Run] listening on %s\n", addr)

	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}
