package main

import (
	"context"
	"flag"
	"net"
	"time"

	"github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver/intercepter"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/creating"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/fetching"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/storage"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/internal/users/updating"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	host := flag.String("host", "localhost:8001", "address to run the server on")
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

	// create repository dependency
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
	usermetaRepo, err := storage.NewMongoClient(conn)
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
