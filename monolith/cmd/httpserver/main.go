package main

import (
	"flag"

	"github.com/KonstantinGasser/datalab/monolith/internal/api"
	"github.com/gorilla/mux"
)

func main() {
	host := flag.String("host", "127.0.0.1:8080", "address to run the server on")
	flag.Parse()

	server := api.New(mux.NewRouter())
	server.Serve(*host)

}
