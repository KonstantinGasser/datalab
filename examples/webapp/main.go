package main

import (
	"log"

	"github.com/KonstantinGasser/clickstream/examples/webapp/cmd/server"
)

/*

 */

func main() {
	// optional:
	// allowing to pass web-server custom server address
	log.Fatal(server.Run(":8081"))
}
