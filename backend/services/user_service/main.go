package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KonstantinGasser/clickstream/backend/services/user_service/cmd/server"
	"github.com/sirupsen/logrus"
)

func main() {
	address := flag.String("listen-addr", ":8001", "address to run the server on")
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
	logrus.Fatal(server.Run(ctx, *address))
}
