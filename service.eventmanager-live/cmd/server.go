package cmd

import (
	"net"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.eventmanager-live/domain"
	"github.com/KonstantinGasser/datalab/service.eventmanager-live/grpcC"
	"github.com/KonstantinGasser/datalab/service.eventmanager-live/handler"
	"github.com/sirupsen/logrus"
)

func Run(host, appTokenAddr, appConfigAddr string) error {

	appTokenSvc, err := grpcC.NewTokenClient(appTokenAddr)
	if err != nil {
		return err
	}
	appConfigSvc, err := grpcC.NewConfigClient(appConfigAddr)
	if err != nil {
		return err
	}

	domain := domain.NewEventLogic(appConfigSvc)
	svc := handler.New(appTokenSvc, domain)
	svc.Apply(
		handler.WithAllowedOrgins("http://localhost:3000", "http://127.0.0.1:3000", "http://192.168.0.232:3000"),
		handler.WithAllowedHeaders("x-datalab-token", "content-type"),
		handler.WithAllowedCreds,
	)

	svc.Register("/api/v1/hello", svc.HandlerInitSession,
		svc.WithCORS,
		svc.WithAuth,
		svc.WithCookie,
	)
	svc.Register("/api/v1/event/live", svc.HandlerOpenSocket,
		svc.WithCORS,
	)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	defer listener.Close()

	logrus.Infof("[api.server] listening on %s\n", host)
	return http.Serve(listener, nil)
}
