package api

import (
	"errors"
	"net/http"
	"strings"

	apptokenSrv "github.com/KonstantinGasser/datalab/protobuf/apptoken_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/utils/hash"
	"github.com/KonstantinGasser/required"
	"github.com/sirupsen/logrus"
)

func (api API) HandlerAppGenerateToken(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppGenerateToken] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)
	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		api.onError(w, errors.New("could not find authenticate for user"), http.StatusUnauthorized)
		return
	}

	var payload struct {
		AppUUID    string `json:"app_uuid" required:"yes"`
		AppName    string `json:"app_name" required:"yes"`
		OrgnDomain string `json:"orgn_domain" required:"yes"`
		OrgnName   string `json:"orgn_name" required:"yes"`
		AppUrl     string `json:"app_url" required:"yes"`
	}
	if err := api.decode(r.Body, &payload); err != nil {
		api.onError(w, errors.New("could not decode request body"), http.StatusBadRequest)
		return
	}
	if err := required.Atomic(&payload); err != nil {
		api.onError(w, errors.New("mandatory fields are missing"), http.StatusBadRequest)
		return
	}
	resp, err := api.AppTokenClient.Issue(r.Context(), &apptokenSrv.IssueRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		CallerUuid: user.GetUuid(),
		AppUuid:    payload.AppUUID,
		AppHash:    hash.Sha256([]byte(strings.Join([]string{payload.OrgnDomain, payload.AppName}, "/"))).String(),
		AppOrigin:  payload.AppUrl,
	})
	if err != nil {
		api.onError(w, errors.New("could not generate App-Token"), http.StatusInternalServerError)
		return
	}
	api.onScucessJSON(w, map[string]interface{}{
		"status":    resp.GetStatusCode(),
		"msg":       resp.GetMsg(),
		"app_token": resp.GetToken(),
	}, int(resp.GetStatusCode()))
}
