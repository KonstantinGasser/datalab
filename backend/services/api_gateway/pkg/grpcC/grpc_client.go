package grpcC

import (
	appSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/app_service"
	tokenSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/token_service"
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

// NewTokenServiceClient is a grpc client
func NewTokenServiceClient(listenOn string) tokenSrv.TokenServiceClient {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil { // change not good !!!
		logrus.Errorf("[NewTokenServiceClient] could not connect: %v", err)
		return nil
	}
	client := tokenSrv.NewTokenServiceClient(conn)
	logrus.Infof("[NewGrpcClient] connected to TokenServiceClient on: %s", listenOn)
	return client
}

// NewAppServiceClient is a grpc client
func NewAppClient(listenOn string) appSrv.AppClient {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil { // change not good !!!
		logrus.Errorf("[NewTokenServiceClient] could not connect: %v", err)
		return nil
	}
	client := appSrv.NewAppClient(conn)
	logrus.Infof("[NewGrpcClient] connected to TokenServiceClient on: %s", listenOn)
	return client
}
