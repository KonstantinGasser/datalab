package main

import (
	"flag"
	"net"

	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/fetching"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/initializing"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/modifying"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/storage/mongo"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	// collect flags
	host := flag.String("host", "localhost:8006", "address to run the server on")
	dbAddr := flag.String("db-srv", "mongodb://apptokenissuer:secure@192.168.0.177:27020", "address to connect to apptoken-database")
	flag.Parse()

	server := grpc.NewServer()
	listener, err := net.Listen("tcp", *host)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("[grpcserver.Listen] listening on: %s\n", *host)
	// create stroage dependency
	apptokenRepo, err := mongo.NewMongoClient(*dbAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	// create service dependencies
	initService := initializing.NewService(apptokenRepo)
	modifyService := modifying.NewService(apptokenRepo)
	fetchService := fetching.NewService(apptokenRepo)

	appTokenServer, err := grpcserver.NewAppTokenServer(initService, modifyService, fetchService)
	if err != nil {
		logrus.Fatal(err)
	}
	proto.RegisterAppTokenServer(server, appTokenServer)

	if err := server.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}
