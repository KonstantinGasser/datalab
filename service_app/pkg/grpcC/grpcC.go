package grpcC

import (
	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	tokenSrv "github.com/KonstantinGasser/datalab/protobuf/token_service"
	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"google.golang.org/grpc"
)

// NewUserClient is a grpc client
func NewUserClient(listenOn string) (userSrv.UserClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := userSrv.NewUserClient(conn)
	return client, nil
}

func NewConfigClient(listenOn string) (configSrv.ConfigClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := configSrv.NewConfigClient(conn)
	return client, nil
}

// NewTokenClient is a grpc client
func NewTokenClient(listenOn string) (tokenSrv.TokenClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := tokenSrv.NewTokenClient(conn)
	return client, nil
}
