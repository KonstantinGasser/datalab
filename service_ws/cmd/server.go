package cmd

import (
	"net"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_ws/pkg/api"
	"github.com/KonstantinGasser/datalab/service_ws/pkg/grpcC"
	"github.com/sirupsen/logrus"
)

func Run(host, tokenAddr, appAddr string) error {
	tokenSrv, err := grpcC.NewTokenClient(tokenAddr)
	if err != nil {
		return err
	}
	appSrv, err := grpcC.NewAppClient(appAddr)
	if err != nil {
		return err
	}

	apisrv := api.New(tokenSrv, appSrv, "")
	apisrv.Apply(
		api.WithAllowedOrgins("http://localhost:3000", "http://127.0.0.1:3000", "http://192.168.0.232:3000"),
		api.WithAllowedHeaders("x-datalab-token", "content-type"),
		api.WithAllowedCreds)

	apisrv.Route("/api/v1/hello",
		apisrv.HandlerInitSession,
		apisrv.WithCORS,
		apisrv.WithAuth,
		apisrv.WithCookie)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	defer listener.Close()

	logrus.Infof("[api.server] listening on %s\n", host)
	return http.Serve(listener, nil)
}
