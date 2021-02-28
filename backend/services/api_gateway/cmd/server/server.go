package server

import (
	"log"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/api_gateway/pkg/api"
)

// Run acts as a run abstraction to start the HTTP-Server
// and can return an error - which is nice when called
// from main
func Run(address string) error {
	srv := api.New(api.CORSConfig{
		Cfgs: []struct {
			Header string
			Value  string
		}{
			{Header: api.AccessControlAllowOrigin, Value: "http://localhost:8081"},
			{Header: api.AccessControlAllowMethods, Value: "POST, OPTIONS"},
			{Header: api.AccessControllAllowHeader, Value: "*"},
		},
	})
	// route and middleware setup
	srv.SetUp()

	log.Printf("listening on %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		return err
	}
	return nil
}
