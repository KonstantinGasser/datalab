package grpc_adapter

import (
	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"google.golang.org/grpc"
)

func NewAppAdministerClient(listenOn string) (appsvc.AppAdministerClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := appsvc.NewAppAdministerClient(conn)
	return client, nil
}
