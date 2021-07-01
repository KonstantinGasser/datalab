package main

import (
	"context"
	"flag"
	"net"
	"time"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/intercepter"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/creating"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/fetching"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/inviting"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/storage"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/internal/apps/updating"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports/client"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "localhost:8003", "address to run the server on")
	appconfigAddr := flag.String("config-srv", "localhost:8005", "address to connect to app-config-service")
	apptokenAddr := flag.String("apptoken-srv", "localhost:8006", "address to connect to user-authentication-service")
	dbAddr := flag.String("db-srv", "mongodb://dev-datalab-user:secure@192.168.0.177:27018", "address to connect to app-database")
	flag.Parse()

	server := grpc.NewServer(
		intercepter.WithUnary(intercepter.WithJwtAuth),
	)
	listener, err := net.Listen("tcp", *host)
	if err != nil {
		logrus.Fatal(err)
	}
	defer listener.Close()
	logrus.Infof("[grpcserver.Listen] listening on: %s\n", *host)

	// create client dependecies
	apptokenEmitter, err := client.NewClientAppToken(*apptokenAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	appConfigEmitter, err := client.NewClientAppConfig(*appconfigAddr)
	if err != nil {
		logrus.Fatal(err)
	}

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
	appmetaRepo, err := storage.NewMongoClient(conn)
	if err != nil {
		logrus.Fatal(err)
	}
	// create service dependencies
	createSerivce := creating.NewService(appmetaRepo, apptokenEmitter, appConfigEmitter)
	fechtService := fetching.NewService(appmetaRepo)
	inviteService := inviting.NewService(appmetaRepo, appConfigEmitter, apptokenEmitter)
	updateService := updating.NewService(appmetaRepo)

	appMetaServer := grpcserver.NewAppMetaServer(
		createSerivce, fechtService, inviteService, updateService)
	proto.RegisterAppMetaServer(server, appMetaServer)

	if err := server.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}
