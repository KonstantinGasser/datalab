package grpcC

import (
	tokenSrv "github.com/KonstantinGasser/datalab/protobuf/token_service"
	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// NewUserClient is a grpc client
func NewUserClient(listenOn string) userSrv.UserClient {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil { // change not good !!!
		logrus.Errorf("[NewUserServiceClient] could not connect: %v", err)
		return nil
	}
	client := userSrv.NewUserClient(conn)
	logrus.Infof("[NewGrpcClient] connected to UserClient on: %s", listenOn)
	return client
}

// NewTokenClient is a grpc client
func NewTokenClient(listenOn string) tokenSrv.TokenClient {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil { // change not good !!!
		logrus.Errorf("[NewTokenClient] could not connect: %v", err)
		return nil
	}
	client := tokenSrv.NewTokenClient(conn)
	logrus.Infof("[NewGrpcClient] connected to TokenClient on: %s", listenOn)
	return client
}
