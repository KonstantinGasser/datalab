package api

import (
	"net/http"

	middelware "github.com/KonstantinGasser/datalab/monolith/internal/api/middleware"
	"github.com/gorilla/mux"
)

type ApiServer struct {
	router *mux.Router
}

func (api ApiServer) Serve(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

func New(router *mux.Router) *ApiServer {
	api := ApiServer{
		router: router,
	}
	api.setUp()
	return &api
}

func (api ApiServer) setUp() {

	api.router.Use(
		middelware.WithLogging,
		middelware.WithIP,
	)
}
