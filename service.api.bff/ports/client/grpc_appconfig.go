package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	grpcAppConfig "github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

type ClientAppConfig struct {
	Conn grpcAppConfig.AppConfigurationClient
}

func NewClientAppConfig(clientAddr string) (*ClientAppConfig, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := grpcAppConfig.NewAppConfigurationClient(conn)
	return &ClientAppConfig{
		Conn: client,
	}, nil
}

func (client ClientAppConfig) CollectAppConfig(ctx context.Context, appUuid string, authedUser *common.AuthedUser, resC chan struct {
	Field string
	Value interface{}
}, errC chan error) {
	resp, err := client.Conn.Get(ctx, &grpcAppConfig.GetRequest{
		Tracing_ID: ctx.Value("tracingID").(string),
		AuthedUser: authedUser,
		AppUuid:    appUuid,
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("appconfig.get err: %s", resp.GetMsg())
		return
	}
	resC <- struct {
		Field string
		Value interface{}
	}{
		Field: "appconfig",
		Value: resp.GetConfigs(),
	}
}
