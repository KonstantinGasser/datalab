package grpc_adapter

import (
	permsvc "github.com/KonstantinGasser/datalab/service.user-permissions/proto"
	"google.golang.org/grpc"
)

func NewUserPermissionsClient(listenOn string) (permsvc.UserPermissionsClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := permsvc.NewUserPermissionsClient(conn)
	return client, nil
}
