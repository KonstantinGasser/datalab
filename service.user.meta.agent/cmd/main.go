package main

import (
	"flag"
	"net"

	"github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/creating"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/fetching"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/storage/mongo"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/updating"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "localhost:8001", "address to run the server on")
	dbAddr := flag.String("db-srv", "mongodb://dev-datalab-user:secure@192.168.0.177:27018", "address to connect to app-database")
	flag.Parse()

	server := grpc.NewServer()
	listener, err := net.Listen("tcp", *host)
	if err != nil {
		logrus.Fatal(err)
	}
	defer listener.Close()
	logrus.Infof("[grpcserver.Listen] listening on: %s\n", *host)

	// create repository dependency
	usermetaRepo, err := mongo.NewMongoClient(*dbAddr)
	if err != nil {
		logrus.Fatal(err)
	}

	// create service dependencies
	createService := creating.NewService(usermetaRepo)
	updateService := updating.NewService(usermetaRepo)
	fetchService := fetching.NewService(usermetaRepo)

	userMetaServer := grpcserver.NewUserMetaServer(
		createService,
		updateService,
		fetchService,
	)
	proto.RegisterUserMetaServer(server, userMetaServer)

	if err := server.Serve(listener); err != nil {
		logrus.Fatal(err)
	}

}
