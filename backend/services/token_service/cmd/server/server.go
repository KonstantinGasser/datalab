package server

import (
	"context"
	"net"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
	"github.com/KonstantinGasser/clickstream/backend/services/token_service/pkg/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, addr string) error {

	srv := grpc.NewServer()

	// create tokenService implementing the grpc methods
	tokenService := api.NewTokenService()

	tokenSrv.RegisterTokenServiceServer(srv, tokenService)

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
