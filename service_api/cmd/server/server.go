package server

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_api/pkg/api"
	"github.com/KonstantinGasser/datalab/service_api/pkg/grpcC"
	"github.com/sirupsen/logrus"
)

// Run acts as a run abstraction to start the HTTP-Server
// and can return an error - which is nice when called
// from main
func Run(ctx context.Context, hostAddr, userAddr, appAddr, tokenAddr, configAddr string) error {
	// create api dependencies
	userClient, err := grpcC.NewUserClient(userAddr)
	if err != nil {
		logrus.Fatalf("[api.Dependency] could not establish gRPC connection to User-Service:\n\t%v", err)
	}
	appClient, err := grpcC.NewAppClient(appAddr)
	if err != nil {
		logrus.Fatalf("[api.Dependency] could not establish gRPC connection to App-Service:\n\t%v", err)
	}
	configClient, err := grpcC.NewConfigClient(configAddr)
	if err != nil {
		logrus.Fatalf("[api.Dependency] could not establish gRPC connection to Config-Service:\n\t%v", err)
	}
	tokenClient, err := grpcC.NewTokenClient(tokenAddr)
	if err != nil {
		logrus.Fatalf("[api.Dependency] could not establish gRPC connection to Token-Service:\n\t%v", err)
	}
	logrus.Info("[api.Dependency] established connection to all services\n")
	srv := api.NewDefault(userClient, appClient, tokenClient, configClient)
	// override default CORS config for Access-Control-Allow-Origin
	// TODO: change "*" to proper value once IP:PORT of frontend is clear
	srv.Apply(srv.WithAccessControlOrigin("*"))

	// route and middleware setup
	srv.SetUp()

	// waiting for context to be canceled
	// not implemented: graceful shutdown
	go func() {
		<-ctx.Done()
		logrus.Infof("Server cleaning up...")
	}()
	logrus.Infof("[server.Run] listening on %s\n", hostAddr)
	if err := http.ListenAndServe(hostAddr, nil); err != nil {
		return err
	}
	return nil
}
