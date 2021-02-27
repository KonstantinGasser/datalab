package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pacedotdev/oto/otohttp"
)

type accountService struct{}

func (accountService) CreateUser(ctx context.Context, r CreateUserRequest) (*CreateUserResponse, error) {
	log.Println(r)
	resp := &CreateUserResponse{
		Status: http.StatusOK,
		Msg:    "first message from oto server! :)",
	}
	return resp, nil
}
func (accountService) LoginUser(ctx context.Context, r LoginUserRequest) (*LoginUserResponse, error) {
	resp := &LoginUserResponse{
		Status: http.StatusOK,
		Msg:    "first message from oto server! :)",
	}
	return resp, nil
}

func run() error {
	otoServer := otohttp.NewServer()
	RegisterAccountService(otoServer, accountService{})
	http.Handle("/oto/", otoServer)

	log.Println("Listening on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		return err
	}
	return nil
}

type CORSHandler struct {
	realServer *otohttp.Server
}

func (c CORSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.WriteHeader(http.StatusOK)
		return
	}
	var d map[string]interface{}
	json.NewDecoder(r.Body).Decode(&d)
	log.Printf("Pre data: %v", d)
	c.realServer.ServeHTTP(w, r)

}
func main() {
	otoServer := otohttp.NewServer()
	RegisterAccountService(otoServer, accountService{})

	c := CORSHandler{realServer: otoServer}
	http.Handle("/oto/", c)

	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-done
		log.Printf("Received SIGNAL: %v\n", sig)
		log.Println("Shutdown of Server...")
		os.Exit(0)

	}()

	log.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
