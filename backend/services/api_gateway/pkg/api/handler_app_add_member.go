package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"

	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// DataAddMember holds the data from the http request
type DataAddMember struct {
	AppUUID string   `json:"app_uuid"`
	Member  []string `json:"member_list"`
}

func (api API) HandlerAppAddMember(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppAddMember] received request\n", ctx_value.GetString(r.Context(), "tracingID"))

	// get authenticated user data to retrieve user uuid
	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		logrus.Errorf("<%v>[api.HandlerAppAddMember] could not find authenticated user\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New("could not find authenticated user"), http.StatusForbidden)
		return
	}

	var payload DataAddMember
	if err := api.decode(r.Body, &payload); err != nil {
		logrus.Errorf("<%v>[api.HandlerAppAddMember] could not decode r.Body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not decode request body"), http.StatusBadRequest)
		return
	}
	// invoke grpc to app service added new user
	resp, err := api.AppServiceClient.AddMember(r.Context(), &appSrv.AddMemberRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		OwnerUuid:  user.GetUuid(),
		AppUuid:    payload.AppUUID,
		Member:     payload.Member,
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppAddMember] could not execute grpc.AddMember: %v\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New("could not add member to app"), http.StatusInternalServerError)
		return
	}
	api.onScucessJSON(w, map[string]interface{}{"status": resp.GetStatusCode(), "msg": resp.GetMsg()}, http.StatusOK)
}