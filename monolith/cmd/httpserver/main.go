package main

import (
	"flag"

	"github.com/KonstantinGasser/datalab/monolith/internal/api"
	userRepo "github.com/KonstantinGasser/datalab/monolith/internal/domain/user/repo"
	userSvc "github.com/KonstantinGasser/datalab/monolith/internal/domain/user/svc"
	"github.com/KonstantinGasser/datalab/monolith/pkg/mongodb"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	host := flag.String("host", "127.0.0.1:8080", "address to run the server on")
	mongoAddr := flag.String("db-srv", "mongodb://dev-datalab-user:secure@192.168.0.178:27018", "address to connect to app-database")
	flag.Parse()

	mongoConn, err := mongodb.NewConn(*mongoAddr)
	if err != nil {
		logrus.Fatal(err)
	}
	// dependencie injection
	// :user-service
	userrepo := userRepo.NewMongoClient(mongoConn)
	usersvc := userSvc.New(userrepo)

	server := api.New(mux.NewRouter(),
		usersvc,
	)
	server.Serve(*host)

}
