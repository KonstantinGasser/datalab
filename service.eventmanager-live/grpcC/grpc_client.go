package grpcC

import (
	appconfig "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	apptokenissuer "github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"google.golang.org/grpc"
)

// NewTokenServiceClient is a grpc client
func NewTokenClient(listenOn string) (apptokenissuer.AppTokenIssuerClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := apptokenissuer.NewAppTokenIssuerClient(conn)
	return client, nil
}

func NewConfigClient(listenOn string) (appconfig.AppConfigurationClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := appconfig.NewAppConfigurationClient(conn)
	return client, nil
}
