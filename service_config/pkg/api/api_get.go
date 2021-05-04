package api

import (
	"context"
	"net/http"

	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv ConfigServer) Get(ctx context.Context, in *configSrv.GetRequest) (*configSrv.GetResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[configService.Get] received request\n", ctx_value.GetString(ctx, "tracingID"))

	cfgItem, err := srv.config.Get(ctx, srv.storage, in.GetConfigUuid())
	if err != nil {
		logrus.Infof("<%v>[configService.Get] could not Get Config: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &configSrv.GetResponse{StatusCode: err.Code(), Msg: err.Info()}, nil
	}

	var stages = make([]*configSrv.Funnel, len(cfgItem.Funnel))
	for i, item := range cfgItem.Funnel {
		stages[i] = &configSrv.Funnel{Id: item.ID, Name: item.Name, Transition: item.Transition}
	}
	var records = make([]*configSrv.Campaign, len(cfgItem.Campaign))
	for i, item := range cfgItem.Campaign {
		records[i] = &configSrv.Campaign{Id: item.ID, Name: item.Name, Prefix: item.Prefix}
	}
	var btnDefs = make([]*configSrv.BtnTime, len(cfgItem.BtnTime))
	for i, item := range cfgItem.BtnTime {
		btnDefs[i] = &configSrv.BtnTime{Id: item.ID, Name: item.Name, BtnName: item.BtnName}
	}
	return &configSrv.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "Configs init done",
		Stages:     stages,
		Records:    records,
		BtnDefs:    btnDefs,
	}, nil
}
