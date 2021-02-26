package main

import (
	"flag"
	"log"

	"github.com/KonstantinGasser/studhouse/backend/api_gateway/cmd/server"
)

func main() {
	address := flag.String("host", ":8080", "address to run the server on")
	flag.Parse()

	log.Fatal(server.Run(*address))

}
