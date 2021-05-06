package api

import (
	"errors"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	apptokenSrv "github.com/KonstantinGasser/datalab/protobuf/apptoken_service"
	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// HandlerAppDetails serves the full JSON required by the frontend for the TAB: APP
// actions required by this TAB can be found under HandlerApp* handlerFuns
func (api API) HandlerAppDetails(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[api.HandlerAppDetails] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	// extract uuid of logged in user
	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		api.onError(w, errors.New("not authenticated"), http.StatusForbidden)
		return
	}

	// fetch all apps mapped to logged in user -> app list in form of: {name, uuid}
	// fetch goes here
	respAppList, err := api.AppClient.GetList(r.Context(), &appSrv.GetListRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		CallerUuid: user.GetUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppDetails] could not execute grpc.GetAppList: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get app details information"), http.StatusInternalServerError)
		return
	}
	// if app list is empty or nil return proper data for user
	if respAppList.GetAppList() == nil || len(respAppList.GetAppList()) == 0 {
		logrus.Infof("<%v>[api.HandlerAppDetails] could not find any apps mapped to caller\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onScucessJSON(w, map[string]interface{}{
			"msg":             "Here you go",
			"app_list":        []struct{}{},
			"app_details":     map[string]interface{}{},
			"app_token":       nil,
			"config_funnel":   nil,
			"config_campaign": nil,
			"config_btn_time": nil,
		}, http.StatusOK)
		return
	}
	// fetch full app details of first app found in fetch of app list (might be changed to latest updated later)
	// fetch goes here
	respAppDetails, err := api.AppClient.Get(r.Context(), &appSrv.GetRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		AppUuid:    respAppList.GetAppList()[0].GetUuid(),
		CallerUuid: user.GetUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppDetails] could not execute grpc.Get (app): %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get app details information"), http.StatusInternalServerError)
		return
	}

	respAppConfig, err := api.ConfigClient.Get(r.Context(), &configSrv.GetRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		ConfigUuid: respAppDetails.GetApp().GetConfigRef(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppDetails] could not execute grpc.Get (config): %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get app details information"), http.StatusInternalServerError)
		return
	}
	respAppToken, err := api.AppTokenClient.Get(r.Context(), &apptokenSrv.GetRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		AppUuid:    respAppList.GetAppList()[0].GetUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[api.HandlerAppDetails] could not execute grpc.Get (apptoken): %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not get app details information"), http.StatusInternalServerError)
		return
	}
	// build response JSON
	// build goes here
	api.onScucessJSON(w, map[string]interface{}{
		"app_list":    respAppList.GetAppList(),
		"app_details": respAppDetails.GetApp(),
		"app_token": map[string]interface{}{
			"token": respAppToken.GetToken(),
			"exp":   respAppToken.GetTokenExp(),
		},
		"config_funnel":   respAppConfig.GetStages(),
		"config_campaign": respAppConfig.GetRecords(),
		"config_btn_time": respAppConfig.GetBtnDefs(),
	}, http.StatusOK)
}
