package main

import (
	"context"
	"flag"
	"net"
	"time"

	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/intercepter"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/fetching"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/initializing"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/modifying"
	configMongo "github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs/storage/mongo"
	libFetch "github.com/KonstantinGasser/datalab/service.app.config.agent/internal/libconfig/fetching"
	libMongo "github.com/KonstantinGasser/datalab/service.app.config.agent/internal/libconfig/storage"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "localhost:8005", "address to run the server on")
	dbAddr := flag.String("db-srv", "mongodb://dev-datalab-user:secure@192.168.0.177:27018", "address to connect to config-database")
	flag.Parse()

	server := grpc.NewServer(
		intercepter.WithUnary(intercepter.WithUserJwtAuth, intercepter.WithSvcJwtAuth),
	)
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
	appconfigRepo := configMongo.NewMongoClient(conn)
	libconfigReqpo := libMongo.NewMongoClient(conn)

	// create service dependencies
	initService := initializing.NewService(appconfigRepo)
	modifyService := modifying.NewService(appconfigRepo)
	fetchService := fetching.NewService(appconfigRepo)

	fetchLibService := libFetch.NewService(libconfigReqpo)

	appConfigServer := grpcserver.NewAppConfigServer(
		initService,
		modifyService,
		fetchService,
		fetchLibService)

	proto.RegisterAppConfigurationServer(server, appConfigServer)

	if err := server.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}
