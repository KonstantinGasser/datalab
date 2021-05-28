package client

import (
	grpcAppConfig "github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	grpcAppToken "github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

type ClientAppToken struct {
	Conn grpcAppToken.AppTokenClient
}

func NewClientAppToken(clientAddr string) (*ClientAppToken, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := grpcAppToken.NewAppTokenClient(conn)
	return &ClientAppToken{
		Conn: client,
	}, nil
}

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
