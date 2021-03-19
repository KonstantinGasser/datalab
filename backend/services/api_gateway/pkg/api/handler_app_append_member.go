package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// DataAppAppendMember represents the data which is required
// in order to add a new member to a given app
type DataAppAppendMember struct {
	Member []string `json:"member_list"`
}

// Todo where is the UUID of the app??? this looks wrong -> pls check on that
// Involved services:
// - App-Service
func (api API) HandlerAppAppendMember(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppAppendMember] received append-app-member request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var payload DataAppAppendMember
	if err := api.decode(r.Body, &payload); err != nil {
		logrus.Errorf("*s*[api.HandlerAppAppendMember] could not decode request body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not decode request body"), http.StatusBadRequest)
	}

	resp, err := api.AppServiceClient.AppendMember(r.Context(), &appSrv.AppendMemberRequest{
		Member: payload.Member,
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppAppendMember] could not execute grpc.AppendMember: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not update app member list"), http.StatusInternalServerError)
		return
	}
	api.onScucessJSON(w, map[string]interface{}{
		"status": resp.GetStatusCode(),
		"msg":    resp.GetMsg(),
	}, int(resp.GetStatusCode()))
}
