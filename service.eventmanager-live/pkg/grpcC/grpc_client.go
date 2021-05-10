package grpcC

import (
	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	tokenSrv "github.com/KonstantinGasser/datalab/protobuf/token_service"
	"google.golang.org/grpc"
)

// NewTokenServiceClient is a grpc client
func NewTokenClient(listenOn string) (tokenSrv.TokenClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := tokenSrv.NewTokenClient(conn)
	return client, nil
}

// NewAppServiceClient is a grpc client
func NewAppClient(listenOn string) (appSrv.AppClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := appSrv.NewAppClient(conn)
	return client, nil
}
