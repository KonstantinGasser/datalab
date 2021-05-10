package grpcC

import (
	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	"google.golang.org/grpc"
)

// NewAppClient is a grpc client
func NewAppClient(listenOn string) (appSrv.AppClient, error) {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := appSrv.NewAppClient(conn)
	return client, nil
}
