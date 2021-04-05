package grpcC

import (
	userSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/user_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewUserServiceClient is a grpc client
func NewUserClient(listenOn string) userSrv.UserClient {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil { // change not good !!!
		logrus.Errorf("[NewUserServiceClient] could not connect: %v", err)
		return nil
	}
	client := userSrv.NewUserClient(conn)
	logrus.Infof("[NewGrpcClient] connected to UserServiceClient on: %s", listenOn)
	return client
}
