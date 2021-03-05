package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/api_gateway/pkg/util"
	"github.com/sirupsen/logrus"
)

func (api API) HandlerGetApps(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerGetApps] received create-app request: %v\n", util.StringValueCtx(r.Context(), "tracingID"), r.Host)

	authedUser := util.AtuhedUserValCtx(r.Context(), "user")
	if authedUser == nil {
		logrus.Errorf("<%v>[api.HandlerGetApps] could not get user-claims, they are nil\n", util.StringValueCtx(r.Context(), "tracingID"))
		api.onError(w, errors.New("could not find user claims"), http.StatusForbidden)
		return
	}
	resp, err := api.AppServiceClient.GetApps(r.Context(), &appSrv.GetAppsRequest{
		UserUuid:   authedUser.GetUuid(),
		Tracing_ID: util.StringValueCtx(r.Context(), "tracingID"),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerGetApps] could not execute grpc.GetApps: %v\n", util.StringValueCtx(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get apps"), http.StatusInternalServerError)
		return
	}
	if resp.GetStatusCode() != 200 {
		logrus.Errorf("<%v>[api.HandlerGetApps] could not get apps from app-service: %v\n", util.StringValueCtx(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get apps"), int(resp.GetStatusCode()))
		return
	}
	api.onScucessJSON(w, resp.GetApps(), http.StatusOK)
}
