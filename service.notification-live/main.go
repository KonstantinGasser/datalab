package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KonstantinGasser/datalab/service.notification-live/cmd/server"
	"github.com/sirupsen/logrus"
)

func main() {
	host := flag.String("host", ":8008", "address to run the server on")
	userauthAddr := flag.String("token-srv", "192.168.0.179:8002", "address to connect to token-service")
	dbAddr := flag.String("db-srv", "mongodb://notify-live-agent:secure@notify-live-agent-db:27022", "address to connect to notification-database")
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
	log.Fatal(server.Run(ctx, *host, *userauthAddr, *dbAddr))
}
