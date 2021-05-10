package server

import (
	"context"
	"net"

	tokenSrv "github.com/KonstantinGasser/datalab/protobuf/token_service"
	"github.com/KonstantinGasser/datalab/service_token/pkg/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Run is a run abstraction for the main function allowing
// to return an error
func Run(ctx context.Context, addr string) error {

	srv := grpc.NewServer()
	// create tokenService implementing the grpc.TokenServiceServer methods
	tokenService := api.NewTokenServer()

	tokenSrv.RegisterTokenServer(srv, tokenService)
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
