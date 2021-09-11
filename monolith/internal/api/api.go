package api

import (
	"net"
	"net/http"

	middelware "github.com/KonstantinGasser/datalab/monolith/internal/api/middleware"
	userSvc "github.com/KonstantinGasser/datalab/monolith/internal/domain/user/svc"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type ApiServer struct {
	router *mux.Router

	// server dependencies
	usersvc userSvc.Service
}

func (api ApiServer) Serve(addr string) {
	listner, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Fatal(err)
	}
	defer listner.Close()
	logrus.Infof("[api-server] serving on %s\n", addr)

	if err := http.Serve(listner, api.router); err != nil {
		logrus.Fatal(err)
	}
}

func New(router *mux.Router, usersvc userSvc.Service) *ApiServer {
	api := ApiServer{
		router:  router,
		usersvc: usersvc,
	}
	api.setUp()
	return &api
}

func (api ApiServer) setUp() {
	api.router.HandleFunc("/api/v1/user/register", api.HandlerRegisterUser).Methods("POST")
	api.router.HandleFunc("/api/v1/user/login", nil).Methods("POST")
	api.router.HandleFunc("/api/v1/user/profile", nil).Methods("POST")
	api.router.HandleFunc("/api/v1/user/profile", nil).Methods("GET")
	api.router.HandleFunc("/api/v1/user/colleagues", nil).Methods("GET")

	api.router.HandleFunc("/api/v1/app/create", nil).Methods("POST")
	api.router.HandleFunc("/api/v1/app", nil).Methods("GET") // also get ALL
	api.router.HandleFunc("/api/v1/app/token/issue", nil).Methods("POST")
	api.router.HandleFunc("/api/v1/app/member", nil).Methods("GET")
	api.router.HandleFunc("/api/v1/app/member/invitable", nil).Methods("GET")

	api.router.HandleFunc("/api/v1/app/config/update", nil).Methods("POST")

	api.router.HandleFunc("/api/v1/app/invite", nil).Methods("POST")
	api.router.HandleFunc("/api/v1/app/invite/accept", nil).Methods("POST")

	api.router.HandleFunc("/api/v1/app/unlock", nil).Methods("POST")

	api.router.Use(
		middelware.WithTracing,
		middelware.WithLogging,
		middelware.WithCors,
	)
}
