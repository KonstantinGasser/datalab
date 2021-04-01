package main

import (
	"flag"
	"log"

	"github.com/KonstantinGasser/clickstream/examples/webapp/cmd/server"
)

/*

 */

func main() {
	// optional:
	// allowing to pass web-server custom server address
	address := flag.String("addr", "localhost:8080", "address the web-server can receive requests")
	log.Fatal(server.Run(*address))
}
