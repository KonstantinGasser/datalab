package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (api API) HandlerAppAppendMember(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppAppendMember] received append-app-member request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)
	data, err := api.decode(r.Body)
	if err != nil {
		logrus.Errorf("*s*[api.HandlerAppAppendMember] could not decode request body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not decode request body"), http.StatusBadRequest)
		return
	}
	// check that passed member list it not nil if use empty slice of string
	var member []string = []string{}
	if _, ok := data["member_list"].([]string); ok {
		member = data["member_list"].([]string)
	}

	resp, err := api.AppServiceClient.AppendMember(r.Context(), &appSrv.AppendMemberRequest{
		Member: member,
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
