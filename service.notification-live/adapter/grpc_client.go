package adapter

import (
	cfgsvc "github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	appsvc "github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	apptokensvc "github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	usersvc "github.com/KonstantinGasser/datalab/service.user-administer/proto"
	userauthsvc "github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

func CreateAppAdminClient(listenon string) (appsvc.AppMetaClient, error) {
	conn, err := grpc.Dial(listenon, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := appsvc.NewAppMetaClient(conn)
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

func CreateAppTokenIssuerClient(listenon string) (apptokensvc.AppTokenClient, error) {
	conn, err := grpc.Dial(listenon, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := apptokensvc.NewAppTokenClient(conn)
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
