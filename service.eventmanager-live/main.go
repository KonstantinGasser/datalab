package main

import (
	"flag"
	"log"

	"github.com/KonstantinGasser/datalab/service.eventmanager-live/cmd"
)

func main() {
	host := flag.String("host", "localhost:8004", "host address to listen on")
	appTokenAddr := flag.String("app-token-srv", "localhost:8006", "address of app-token service")
	appConfigAddr := flag.String("app-config-srv", "localhost:8005", "address of app-config service")
	// appAddr := flag.String("app-srv", "localhost:8003", "address of app service")
	flag.Parse()

	log.Fatal(cmd.Run(*host, *appTokenAddr, *appConfigAddr))
}
