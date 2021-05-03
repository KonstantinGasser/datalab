package cmd

import (
	"fmt"
	"net"
	"net/http"

	"github.com/KonstantinGasser/datalab/service_ws/pkg/api"
)

func Run() error {

	apisrv := api.New("does not matter rn")
	apisrv.Route("/cookie", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello friend - got a cookie?")
	}, apisrv.WithCookie)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}
	return http.Serve(listener, nil)
}
