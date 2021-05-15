package grpc_adapter

import (
	cfgsvc "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	aptissuer "github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	usersvc "github.com/KonstantinGasser/datalab/service.user-administer/proto"
	userauthsvc "github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"google.golang.org/grpc"
)

// NewUserAdministerClient is a grpc client
func NewUserAdministerClient(listenOn string) (usersvc.UserAdministerClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := usersvc.NewUserAdministerClient(conn)
	return client, nil
}

func NewAppConfigClient(listenOn string) (cfgsvc.AppConfigurationClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := cfgsvc.NewAppConfigurationClient(conn)
	return client, nil
}

// NewAppTokenIssuerClient is a grpc client
func NewAppTokenIssuerClient(listenOn string) (aptissuer.AppTokenIssuerClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := aptissuer.NewAppTokenIssuerClient(conn)
	return client, nil
}

func NewUserAuthClient(listenOn string) (userauthsvc.UserAuthenticationClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := userauthsvc.NewUserAuthenticationClient(conn)
	return client, nil
}
