package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KonstantinGasser/datalab/service_api/cmd/server"
	"github.com/sirupsen/logrus"
)

func main() {
	host := flag.String("host", "localhost:8080", "address to run the server on")
	userAddr := flag.String("user-srv", "localhost:8001", "address to connect to user-service")
	appAddr := flag.String("app-srv", "localhost:8003", "address to connect to app-service")
	configAddr := flag.String("config-srv", "localhost:8005", "address to connect to app-service")
	tokenAddr := flag.String("token-srv", "localhost:8002", "address to connect to token-service")
	flag.Parse()

	// SIG chan to handle interruptions and so on...
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, os.Kill)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sig := <-done
		logrus.Warnf("Received OS signal - shutting down... SIG: %s\n", sig)
		cancel()
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()
	// call Run abstraction to start the server
	log.Fatal(server.Run(ctx, *host, *userAddr, *appAddr, *tokenAddr, *configAddr))
}
