package main

import (
	"flag"
	"log"

	"github.com/KonstantinGasser/clickstream/backend/services/api_gateway/cmd/server"
)

func main() {
	address := flag.String("listen-addr", "localhost:8080", "address to run the server on")
	flag.Parse()

	log.Fatal(server.Run(*address))
}
