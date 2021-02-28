package grpcC

import (
	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewUserServiceClient(listenOn string) userSrv.UserServiceClient {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil { // change not good !!!
		logrus.Errorf("[NewUserServiceClient] could not connect: %v", err)
		return nil
	}
	client := userSrv.NewUserServiceClient(conn)
	return client
}
