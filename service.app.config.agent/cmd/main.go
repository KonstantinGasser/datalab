package main

import (
	"flag"
	"net"

	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/intercepter"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/fetching"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/initializing"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/modifying"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/storage/mongo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "localhost:8005", "address to run the server on")
	dbAddr := flag.String("db-srv", "mongodb://dev-datalab-user:secure@192.168.0.177:27018", "address to connect to config-database")
	flag.Parse()

	server := grpc.NewServer(
		intercepter.WithUnary(intercepter.WithJwtAuth),
	)
	listener, err := net.Listen("tcp", *host)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("[grpcserver.Listen] listening on: %s\n", *host)

	// create storage dependency
	appconfigRepo, err := mongo.NewMongoClient(*dbAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	// create service dependencies
	initService := initializing.NewService(appconfigRepo)
	modifyService := modifying.NewService(appconfigRepo)
	fetchService := fetching.NewService(appconfigRepo)

	appConfigServer := grpcserver.NewAppConfigServer(initService, modifyService, fetchService)
	proto.RegisterAppConfigurationServer(server, appConfigServer)

	if err := server.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}
