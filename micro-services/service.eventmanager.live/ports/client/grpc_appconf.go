package client

import (
	"context"
	"net/http"

	grpcAppConfig "github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.eventmanager.live/ports/client/intercepter"
	"google.golang.org/grpc"
)

type ClientAppConfig struct {
	Conn grpcAppConfig.AppConfigurationClient
}

func NewClientAppConfig(clientAddr string) (*ClientAppConfig, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure(),
		intercepter.WithUnary(intercepter.WithServiceAuth))
	if err != nil {
		return nil, err
	}
	client := grpcAppConfig.NewAppConfigurationClient(conn)
	return &ClientAppConfig{
		Conn: client,
	}, nil
}

func (client ClientAppConfig) GetAppConfig(ctx context.Context, appUuid string) (*grpcAppConfig.ClientConfig, error) {
	resp, err := client.Conn.GetForClient(ctx, &grpcAppConfig.GetForClientRequest{
		AppUuid: appUuid,
	})
	if err != nil {
		return nil, err
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, err
	}
	return resp.GetConfig(), nil
}
