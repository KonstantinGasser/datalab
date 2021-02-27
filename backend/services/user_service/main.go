package main

import (
	"flag"
	"log"

	"github.com/KonstantinGasser/clickstream/backend/services/user_service/cmd/server"
)

func main() {
	address := flag.String("listen-addr", ":8001", "address to run the server on")
	flag.Parse()

	log.Fatal(server.Run(*address))
}
