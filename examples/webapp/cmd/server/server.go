package server

import (
	"log"
	"net"
	"net/http"

	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/api"
	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/services/user"
	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/storage"
)

// Run is a abstraction for the main function allowing to return an error
// so it can be handeld in the main func
func Run(serverAddress string) error {
	// create network listener listening on TCP:somePort
	listener, err := net.Listen("tcp", serverAddress)
	defer listener.Close()
	if err != nil {
		log.Fatalf("could not start listener on: %s:%v", serverAddress, err)
		return err
	}

	// create new user-service as api dep
	userSrv := user.New()
	// create new storage as api dep
	storage := storage.New("in-memory")
	// create API instance
	api := api.New(userSrv, storage)
	// setup routes and init API
	api.SetUp()

	log.Printf("starting HTTP-Server on: %s\n", serverAddress)
	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("could not start http Server: %v", err)
		return err
	}
	return nil
}