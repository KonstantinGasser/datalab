package server

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/backend/services/file_server/pkg/api"
	"github.com/sirupsen/logrus"
)

// Run acts as a run abstraction to start the HTTP-Server
// and can return an error - which is nice when called
// from main
func Run(ctx context.Context, address string) error {

	srv := api.New()
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
