package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	grpcAppConfig "github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports/client/intercepter"
	"google.golang.org/grpc"
)

type ClientAppConfig struct {
	conn grpcAppConfig.AppConfigurationClient
}

func NewClientAppConfig(clientAddr string) (*ClientAppConfig, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure(), intercepter.WithUnary(intercepter.WithAuth))
	if err != nil {
		return nil, err
	}
	client := grpcAppConfig.NewAppConfigurationClient(conn)
	return &ClientAppConfig{
		conn: client,
	}, nil
}

// Emit tirggers the initialize endpoint for the AppConfigSerivce asking it to
// initialize and acknowlege that a new app config with the given reference needs to be
// created
func (client ClientAppConfig) Emit(ctx context.Context, event *ports.Event, errC chan error) {
	resp, err := client.conn.Initialize(ctx, &grpcAppConfig.InitRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppRefUuid: event.App.Uuid,
		AppOwner:   event.App.OwnerUuid,
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("appconf.init err: %s", resp.GetMsg())
		return
	}
	errC <- nil
}
