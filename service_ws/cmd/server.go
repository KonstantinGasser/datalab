package cmd

import (
	"net"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_ws/pkg/api"
	"github.com/sirupsen/logrus"
)

func Run(listenOn string) error {
	apisrv := api.New("does not matter rn")

	apisrv.Apply(
		api.WithAllowedOrgins("http://localhost:3000", "http://127.0.0.1:3000", "http://192.168.0.232:3000"),
		api.WithAllowedHeaders("x-datalab-token", "content-type"),
		api.WithAllowedCreds)

	apisrv.Route("/api/v1/hello",
		apisrv.HandlerInitSession,
		apisrv.WithCORS,
		apisrv.WithAuth,
		apisrv.WithCookie)

	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		return err
	}
	logrus.Infof("[api.server] listening on %s\n", listenOn)
	return http.Serve(listener, nil)
}
