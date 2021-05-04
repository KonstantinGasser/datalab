package main

import (
	"flag"
	"log"

	"github.com/KonstantinGasser/datalab/service_ws/cmd"
)

func main() {
	host := flag.String("host", "192.168.0.232:8004", "host address to listen on")
	flag.Parse()

	log.Fatal(cmd.Run(*host))
}
