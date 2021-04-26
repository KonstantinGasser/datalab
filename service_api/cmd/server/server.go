package server

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalabs/service_api/pkg/api"
	"github.com/sirupsen/logrus"
)

// Run acts as a run abstraction to start the HTTP-Server
// and can return an error - which is nice when called
// from main
func Run(ctx context.Context, address string) error {
	// creating new api with configs for CORS settings
	// due to the pre-flight call via OPTIONS from the browser
	// setting of allowed-origin, allowed-methods, and allowed-headers
	// is required
	srv := api.New(api.CORSConfig{
		Cfgs: []struct {
			Header string
			Value  string
		}{
			{Header: api.AccessControlAllowOrigin, Value: "http://localhost:3000"},
			{Header: api.AccessControlAllowMethods, Value: "GET,POST, OPTIONS"},
			{Header: api.AccessControlAllowHeader, Value: "*"},
		},
	})
	// route and middleware setup
	srv.SetUp()

	// waiting for context to be canceled
	// not implemented: graceful shutdown
	go func() {
		<-ctx.Done()
		logrus.Infof("Server cleaning up...")
	}()
	logrus.Infof("[server.Run] listening on %s\n", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		return err
	}
	return nil
}
