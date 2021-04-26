package api

import (
	"fmt"
	"net/http"

	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// HandlerUserRegister is the entry-point if a users creates a new account.
// It performs sanity checks on the input data and forwards the request
// to the user-service.
// Involved services:
//	- User-Service
func (api API) HandlerUserRegister(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerRegister] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var payload struct {
		Username     string `json:"username"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Password     string `json:"password"`
		OrgnDomain   string `json:"orgn_domain"`
		OrgnPosition string `json:"orgn_position"`
	}
	if err := api.decode(r.Body, &payload); err != nil {
		api.onError(w, err, http.StatusBadRequest)
		return
	}
	// passed data is not allowed to be empty
	// Todo: create helper func on api to perform checks on N inputs
	if payload.Username == "" || payload.Password == "" || payload.OrgnDomain == "" {
		api.onError(w, fmt.Errorf("missing fields in register request"), http.StatusBadRequest)
		return
	}

	// invoke grpc call to user-service to create the user
	// Response holds only a status-code and a msg (could be an error message)
	resp, err := api.UserClient.Create(r.Context(), &userSrv.CreateRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		User: &userSrv.RegisterUser{
			Username:     payload.Username,
			Password:     payload.Password,
			FirstName:    payload.FirstName,
			LastName:     payload.LastName,
			OrgnDomain:   payload.OrgnDomain,
			OrgnPosition: payload.OrgnPosition,
		},
	})
	if err != nil {
		api.onError(w, fmt.Errorf("could not execute grpc.CreateUser: %v", err), http.StatusInternalServerError)
		return
	}
	// on success write to response
	api.onScucessJSON(w, map[string]interface{}{}, int(resp.GetStatusCode()))
}
