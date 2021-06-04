package main

import (
	"flag"
	"net"

	"github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/adding"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/fetching"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/initializing"
	mongoPermissions "github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/storage/mongo"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/users/authenticating"
	mongoUser "github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/users/storage/mongo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "localhost:8002", "address to run the server on")
	dbAddr := flag.String("db-srv", "mongodb://userauthstorage:secure@192.168.178.103:27021", "address to connect to user-auth-database")
	flag.Parse()

	server := grpc.NewServer()
	listener, err := net.Listen("tcp", *host)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("[grpcserver.Listen] listening on: %s\n", *host)

	// create storage dependency
	userRepo, err := mongoUser.NewMongoClient(*dbAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	permissionRepo, err := mongoPermissions.NewMongoClient(*dbAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	// create service dependencie
	initService := initializing.NewService(permissionRepo)

	authService := authenticating.NewService(userRepo, initService)
	addService := adding.NewService(permissionRepo)
	fetchService := fetching.NewService(permissionRepo)

	userauthServer := grpcserver.NewUserAuthServer(authService, addService, fetchService)
	proto.RegisterUserAuthenticationServer(server, userauthServer)

	if err := server.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}
