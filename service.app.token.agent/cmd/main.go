package main

import (
	"context"
	"flag"
	"net"
	"time"

	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/intercepter"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/fetching"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/initializing"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/modifying"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/storage"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/internal/apptokens/validating"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {

	// collect flags
	host := flag.String("host", "localhost:8006", "address to run the server on")
	dbAddr := flag.String("db-srv", "mongodb://dev-datalab-user:secure@192.168.0.177:27018", "address to connect to apptoken-database")
	flag.Parse()

	server := grpc.NewServer(
		intercepter.WithUnary(intercepter.WithJwtAuth),
	)
	listener, err := net.Listen("tcp", *host)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("[grpcserver.Listen] listening on: %s\n", *host)
	// create stroage dependency
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
	apptokenRepo, err := storage.NewMongoClient(conn)
	if err != nil {
		logrus.Fatal(err)
	}
	// create service dependencies
	initService := initializing.NewService(apptokenRepo)
	modifyService := modifying.NewService(apptokenRepo)
	fetchService := fetching.NewService(apptokenRepo)
	validateServcie := validating.NewService(apptokenRepo)

	appTokenServer, err := grpcserver.NewAppTokenServer(initService, modifyService, fetchService, validateServcie)
	if err != nil {
		logrus.Fatal(err)
	}
	proto.RegisterAppTokenServer(server, appTokenServer)

	if err := server.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}
