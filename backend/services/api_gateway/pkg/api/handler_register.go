package api

import (
	"context"
	"fmt"
	"net/http"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/sirupsen/logrus"
)

// HandlerRegister receives request to create a new user
func (api API) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	// get data from r.Body
	data, err := api.decode(r.Body)
	if err != nil {
		api.onError(w, err, http.StatusBadRequest)
		return
	}
	// check posted data for correctness
	if data["username"].(string) == "" || data["password"] == "" || data["orgn_domain"] == "" {
		api.onError(w, fmt.Errorf("missing fields in register request"), http.StatusBadRequest)
		return
	}
	// invoke grpc call
	resp, err := api.UserSrvClient.CreateUser(context.Background(), &userSrv.CreateUserRequest{
		Username:   data["username"].(string),
		Password:   data["password"].(string),
		OrgnDomain: data["orgn_domain"].(string),
	})
	if err != nil {
		api.onError(w, fmt.Errorf("could not reach UserService for request: %v", err), http.StatusInternalServerError)
		return
	}
	logrus.Infof("[grpc.CreateUser] status: %d, msg: %s", resp.GetStatusCode(), resp.GetMsg())
	// return success of register request

	b, err := api.encode(resp)
	if err != nil {
		logrus.Errorf("[api.HandlerRegister] could not encode grpc response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(resp.GetStatusCode()))
	w.Write(b)
}
