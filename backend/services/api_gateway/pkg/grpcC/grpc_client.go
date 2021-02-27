package grpcC

import (
	"log"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"google.golang.org/grpc"
)

func NewUserServiceClient(listenOn string) userSrv.UserServiceClient {
	conn, err := grpc.Dial(listenOn, grpc.WithInsecure())
	if err != nil { // change not good !!!
		log.Fatal(err)
		return nil
	}
	client := userSrv.NewUserServiceClient(conn)
	return client
}
