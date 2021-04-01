package cmd

import (
	"log"
	"net"
	"net/http"
)

func Run(serverAddress string) error {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("could not start listener on: %s:%v", serverAddress, err)
		return err
	}

	log.Printf("starting HTTP-Server on: %s\n", serverAddress)
	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("could not start http Server: %v", err)
		return err
	}
	return nil
}
