package api

import (
	"context"
	"log"
	"net/http"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
)

// HandlerLogin forwards the login request to the user and token service
// in oder to check the user's auth and to issue a JWT.
func (api API) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// invoke grpc call
	resp, err := api.UserSrvClient.AuthUser(context.Background(), &userSrv.AuthRequest{
		Username: "KonstantinGasser",
		Password: "secure-hash",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("First grpc response yeay: %v\n", resp)
	w.WriteHeader(http.StatusOK)
	return
}
