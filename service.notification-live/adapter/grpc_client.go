package adapter

import (
	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	cfgsvc "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	apptokensvc "github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	usersvc "github.com/KonstantinGasser/datalab/service.user-administer/proto"
	userauthsvc "github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"google.golang.org/grpc"
)

func CreateAppAdminClient(listenon string) (appsvc.AppAdministerClient, error) {
	conn, err := grpc.Dial(listenon, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := appsvc.NewAppAdministerClient(conn)
	return client, nil
}
func CreateAppConfigClient(listenon string) (cfgsvc.AppConfigurationClient, error) {
	conn, err := grpc.Dial(listenon, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := cfgsvc.NewAppConfigurationClient(conn)
	return client, nil
}

func CreateAppTokenIssuerClient(listenon string) (apptokensvc.AppTokenIssuerClient, error) {
	conn, err := grpc.Dial(listenon, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := apptokensvc.NewAppTokenIssuerClient(conn)
	return client, nil
}

func CreateUserAdminClient(listenon string) (usersvc.UserAdministerClient, error) {
	conn, err := grpc.Dial(listenon, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := usersvc.NewUserAdministerClient(conn)
	return client, nil
}
func CreateUserAuthClient(listenon string) (userauthsvc.UserAuthenticationClient, error) {
	conn, err := grpc.Dial(listenon, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := userauthsvc.NewUserAuthenticationClient(conn)
	return client, nil
}
