package server

import (
	"log"
	"net"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/backend/services/user_service/pkg/api"
	grpc "google.golang.org/grpc"
)

// Run is a run-abstraction for the main func
func Run(addr string) error {
	srv := grpc.NewServer()
	userSrv.RegisterUserServiceServer(srv, api.UserService{})
	// create tcp listener
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("cloud not listen to %s: %v", addr, err)
	}
	log.Printf("listening on %s...\n", addr)

	if err := srv.Serve(listener); err != nil {
		return err
	}
	return nil
}
