package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/api_gateway/pkg/util"
	"github.com/sirupsen/logrus"
)

// HandlerAppCreate is the api endpoint if a logged in user wants to create a new application
func (api API) HandlerAppCreate(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppCreate] received create-app request: %v\n", util.StringValueCtx(r.Context(), "tracingID"), r.Host)

	payload, err := api.decode(r.Body)
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppCreate] could not decode r.Body: %v\n", util.StringValueCtx(r.Context(), "tracingID"), err)
	}
	authedUser := util.AtuhedUserValCtx(r.Context(), "user")
	if authedUser == nil {
		logrus.Errorf("<%v>[api.HandlerAppCreate] could not get user-claims, they are nil\n", util.StringValueCtx(r.Context(), "tracingID"))
		api.onError(w, errors.New("could not find user claims"), http.StatusForbidden)
		return
	}
	respApp, err := api.AppServiceClient.CreateApp(r.Context(), &appSrv.CreateAppRequest{
		OwnerUuid:    authedUser.GetUuid(),
		Name:         payload["app_name"].(string),
		Organization: authedUser.GetOrgnDomain(),
		Tracing_ID:   util.StringValueCtx(r.Context(), "tracingID"),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppCreate] could not execute grpc.CreateApp: %v\n", util.StringValueCtx(r.Context(), "tracingID"), err)
		api.onError(w, err, http.StatusInternalServerError)
		return
	}
	logrus.Info(respApp)
	api.onScucessJSON(w, `{"msg": "worked!"}`, http.StatusOK)
}
