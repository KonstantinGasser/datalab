package grpcC

import (
	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewUserServiceClient is a grpc client
func NewUserServiceClient(listenOn string) userSrv.UserServiceClient {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil { // change not good !!!
		logrus.Errorf("[NewUserServiceClient] could not connect: %v", err)
		return nil
	}
	client := userSrv.NewUserServiceClient(conn)
	logrus.Infof("[NewGrpcClient] connected to UserServiceClient on: %s", listenOn)
	return client
}

// NewTokenServiceClient is a grpc client
func NewTokenServiceClient(listenOn string) tokenSrv.TokenServiceClient {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil {
		logrus.Errorf("[NewTokenServiceClient] could not connect: %v", err)
		return nil
	}
	client := tokenSrv.NewTokenServiceClient(conn)
	logrus.Infof("[NewGrpcClient] connected to TokenServiceClient on: %s", listenOn)
	return client
}
