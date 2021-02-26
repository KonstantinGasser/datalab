package server

import (
	"log"
	"net/http"

	"github.com/KonstantinGasser/studhouse/backend/api_gateway/pkg/api"
)

// Run acts as a run abstraction to start the HTTP-Server
// and can return an error - which is nice when called
// from main
func Run(address string) error {
	srv := api.NewAPI(api.Config{
		AccessControlAllowOrigin:  "*",
		AccessControlAllowMethods: "POST, OPTIONS",
	})

	// setup routes
	srv.Route("/login", srv.WithCors(srv.HandlerLogin))

	log.Printf("listening on %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		return err
	}
	return nil
}
