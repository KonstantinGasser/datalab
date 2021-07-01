package main

import (
	"context"
	"flag"
	"net"
	"time"

	"github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/adding"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/fetching"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/initializing"
	mongoPermissions "github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/permissions/storage"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/users/authenticating"
	mongoUser "github.com/KonstantinGasser/datalab/service.user.auth.agent/internal/users/storage"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "localhost:8002", "address to run the server on")
	dbAddr := flag.String("db-srv", "mongodb://dev-datalab-user:secure@192.168.0.177:27018", "address to connect to user-auth-database")
	flag.Parse()

	server := grpc.NewServer()
	listener, err := net.Listen("tcp", *host)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("[grpcserver.Listen] listening on: %s\n", *host)

	// create storage dependency
	opts := options.Client().ApplyURI(*dbAddr)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	conn, err := mongo.Connect(ctx, opts)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := conn.Ping(context.TODO(), nil); err != nil {
		logrus.Fatal(err)
	}
	userRepo, err := mongoUser.NewMongoClient(conn)
	if err != nil {
		logrus.Fatal(err)
	}
	permissionRepo, err := mongoPermissions.NewMongoClient(conn)
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
