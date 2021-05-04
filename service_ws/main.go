package main

import (
	"flag"
	"log"

	"github.com/KonstantinGasser/datalab/service_ws/cmd"
)

func main() {
	host := flag.String("host", "192.168.0.232:8004", "host address to listen on")
	tokenAddr := flag.String("token-srv", "localhost:8002", "address of token service")
	appAddr := flag.String("app-srv", "localhost:8003", "address of app service")
	flag.Parse()

	log.Fatal(cmd.Run(*host, *tokenAddr, *appAddr))
}
