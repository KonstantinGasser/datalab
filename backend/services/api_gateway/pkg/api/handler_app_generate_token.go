package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/app_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
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
		AppUUID  string `json:"app_uuid"`
		AppName  string `json:"app_name"`
		OrgnName string `json:"orgn_name"`
	}
	if err := api.decode(r.Body, &payload); err != nil {
		api.onError(w, errors.New("could not decode request body"), http.StatusBadRequest)
		return
	}
	resp, err := api.AppClient.GenerateToken(r.Context(), &appSrv.GenerateTokenRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		CallerUuid: user.GetUuid(),
		AppUuid:    payload.AppUUID,
		AppName:    payload.AppName,
		OrgnDomain: payload.OrgnName,
	})
	if err != nil {
		api.onError(w, errors.New("could not generate app token"), http.StatusInternalServerError)
		return
	}
	api.onScucessJSON(w, map[string]interface{}{
		"status":    resp.GetStatusCode(),
		"msg":       resp.GetMsg(),
		"app_token": resp.GetAppToken(),
	}, int(resp.GetStatusCode()))
}
