package api

import (
	"encoding/json"
	"errors"
	"net/http"

	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/required"
	"github.com/sirupsen/logrus"
)

func (api API) HandlerAppUpdateConfig(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%s>[api.HandlerAppUpdateConfig] received request\n", ctx_value.GetString(r.Context(), "tracingID"))

	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		logrus.Errorf("<%s>[api.HandlerAppUpdateConfig] user not authenticated: %v\n", ctx_value.GetString(r.Context(), "tracingID"))
		api.onError(w, errors.New("request not authenticated"), http.StatusUnauthorized)
		return
	}

	var payload struct {
		ConfigUUID string `json:"config_uuid" required:"yes"`
		Funnel     []struct {
			ID         int32  `json:"id"`
			Name       string `json:"name"`
			Transition string `json:"transition"`
		} `json:"funnel"`
		Campaign []struct {
			ID     int32  `json:"id"`
			Name   string `json:"name"`
			Prefix string `json:"prefix"`
		} `json:"campaign"`
		BtnTime []struct {
			ID      int32  `json:"id"`
			Name    string `json:"name"`
			BtnName string `json:"btn_name"`
		} `json:"btn_time"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		logrus.Errorf("<%s>[api.HandlerAppUpdateConfig] could not decode r.Body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not decode request body"), http.StatusBadRequest)
		return
	}
	if err := required.Atomic(&payload); err != nil {
		logrus.Errorf("<%s>[api.HandlerAppUpdateConfig] mandatory fields in r.Body missing: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("request missing mandatory fields"), http.StatusBadRequest)
		return
	}

	// updateFlag can be "funnel", "campaign" or "btn_time" do we need "*" as wildcard??(update all)
	updateFlag := api.getQuery(r.URL, "resource")

	var stages = make([]*configSrv.Funnel, len(payload.Funnel))
	for i, item := range payload.Funnel {
		stages[i] = &configSrv.Funnel{Id: item.ID, Name: item.Name, Transition: item.Transition}
	}
	var records = make([]*configSrv.Campaign, len(payload.Campaign))
	for i, item := range payload.Campaign {
		records[i] = &configSrv.Campaign{Id: item.ID, Name: item.Name, Prefix: item.Prefix}
	}
	var btnDefs = make([]*configSrv.BtnTime, len(payload.BtnTime))
	for i, item := range payload.BtnTime {
		btnDefs[i] = &configSrv.BtnTime{Id: item.ID, Name: item.Name, BtnName: item.BtnName}
	}
	resp, err := api.ConfigClient.Update(r.Context(), &configSrv.UpdateRequest{
		Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		UUID:       payload.ConfigUUID,
		UpdateFlag: updateFlag,
		Stages:     stages,
		Records:    records,
		BtnDefs:    btnDefs,
	})
	if err != nil {
		logrus.Errorf("<%s>[api.HandlerAppUpdateConfig] could not execute grpc.Update call: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		api.onError(w, errors.New("could not change app configurations"), http.StatusInternalServerError)
		return
	}
	// resp, err := api.AppClient.UpdateCfg(r.Context(), &appSrv.UpdateCfgRequest{
	// 	Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
	// 	CallerUuid: user.GetUuid(),
	// 	UpdateFlag: updateFlag,
	// 	AppUuid:    payload.AppUUID,
	// 	Stages:     funnelStages,
	// 	Records:    campaignRecords,
	// 	BtnDefs:    btnDefs,
	// })
	// if err != nil {
	// 	logrus.Errorf("<%s>[api.HandlerAppUpdateConfig] could not execute grpc call: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
	// 	api.onError(w, errors.New("could not change app configurations"), http.StatusInternalServerError)
	// 	return
	// }
	api.onScucessJSON(w, map[string]interface{}{}, int(resp.GetStatusCode()))
}
