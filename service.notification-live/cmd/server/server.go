package server

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.notification-live/adapter"
	"github.com/KonstantinGasser/datalab/service.notification-live/domain"
	"github.com/KonstantinGasser/datalab/service.notification-live/handler"
	"github.com/KonstantinGasser/datalab/service.notification-live/notifyhub"
	"github.com/sirupsen/logrus"
)

// Run acts as a run abstraction to start the HTTP-Server
// and can return an error - which is nice when called
// from main
func Run(ctx context.Context, hostAddr, userauthAddr string) error {
	// create api dependencies
	userauthClient, err := adapter.CreateUserAuthClient(userauthAddr)
	if err != nil {
		return err
	}

	notifyHub := notifyhub.New()
	defer notifyHub.Stop()

	domain := domain.NewNotificationLogic(userauthClient, notifyHub)

	notifySvc := handler.NewHandler(domain)
	logrus.Info("[api.Dependency] established connection to all services\n")
	notifySvc.Apply(handler.WithAllowedOrigins("*"))

	notifySvc.Register("/api/v1/datalab/live", notifySvc.HandlerOpenSocket,
		notifySvc.WithTracing,
		notifySvc.WithAuth,
		notifySvc.WithCors,
	)

	notifySvc.Register("/api/v1/datalab/publish/event", notifySvc.HandleIncomingNofication,
		notifySvc.WithTracing,
		notifySvc.WithCors,
	)
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
