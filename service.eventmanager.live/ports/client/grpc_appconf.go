package client

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
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

func (client ClientAppConfig) GetAppConfig(ctx context.Context, appUuid string, authedUser *common.AuthedUser) (*common.AppConfigurations, error) {
	resp, err := client.Conn.Get(ctx, &grpcAppConfig.GetRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AuthedUser: authedUser,
		AppUuid:    appUuid,
	})
	if err != nil {
		return nil, err
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, err
	}
	return resp.GetConfigs(), nil
}
