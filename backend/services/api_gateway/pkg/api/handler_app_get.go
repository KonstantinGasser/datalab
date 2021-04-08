package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/app_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (api API) HandlerAppGet(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppGet] received request\n", ctx_value.GetString(r.Context(), "tracingID"))

	appUUID := api.getQuery(r.URL, "uuid")
	if appUUID == "" {
		logrus.Errorf("<%v>[api.HandlerAppGet] could not find app uuid in URL query\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New("could not find any app-uuid in query"), http.StatusBadRequest)
		return
	}

	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		logrus.Errorf("<%v>[api.HandlerAppGet] could not find authenticated user\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New("user not authenticated"), http.StatusUnauthorized)
		return
	}
	resp, err := api.AppClient.Get(r.Context(), &appSrv.GetRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		AppUuid:    appUUID,
		CallerUuid: user.GetUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppGet] could not execute grpc.Get: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get app details"), http.StatusInternalServerError)
		return
	}
	logrus.Warn(resp.GetApp())
	api.onScucessJSON(w, map[string]interface{}{
		"status": resp.GetStatusCode(),
		"msg":    resp.GetMsg(),
		"app":    resp.GetApp(),
	}, http.StatusOK)
}
