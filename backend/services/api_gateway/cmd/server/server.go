package server

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/clickstream/backend/services/api_gateway/pkg/api"
	"github.com/sirupsen/logrus"
)

// Run acts as a run abstraction to start the HTTP-Server
// and can return an error - which is nice when called
// from main
func Run(ctx context.Context, address string) error {
	srv := api.New(api.CORSConfig{
		Cfgs: []struct {
			Header string
			Value  string
		}{
			{Header: api.AccessControlAllowOrigin, Value: "http://localhost:3000"},
			{Header: api.AccessControlAllowMethods, Value: "POST, OPTIONS"},
			{Header: api.AccessControllAllowHeader, Value: "*"},
		},
	})
	// route and middleware setup
	srv.SetUp()

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
