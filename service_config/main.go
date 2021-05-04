package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KonstantinGasser/datalab/service_config/cmd/server"
	"github.com/sirupsen/logrus"
)

func main() {
	host := flag.String("host", "localhost:8005", "address to run the server on")
	dbAddr := flag.String("db-srv", "mongodb://configstorage:secure@192.168.0.177:27019", "address to connect to config-database")
	flag.Parse()

	// SIG chan to handle interruptions and so on...
	ctx, cancle := context.WithCancel(context.Background())
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, os.Kill, syscall.SIGTERM)
	go func() {
		sig := <-done
		logrus.Warnf("Received OS signal - shutting down... SIG: %s\n", sig)
		cancle()
		time.Sleep(time.Second * 1)
		os.Exit(0)
	}()
	logrus.Fatal(server.Run(ctx, *host, *dbAddr))
}
