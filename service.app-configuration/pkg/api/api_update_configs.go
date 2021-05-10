package api

import (
	"context"
	"net/http"

	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	"github.com/KonstantinGasser/datalab/service_config/pkg/config"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv ConfigServer) Update(ctx context.Context, in *configSrv.UpdateRequest) (*configSrv.UpdateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[configService.UpdateConfigs] received request\n", ctx_value.GetString(ctx, "tracingID"))

	var stages = make([]config.Stage, len(in.GetStages()))
	for i, item := range in.GetStages() {
		stages[i] = config.Stage{ID: item.Id, Name: item.Name, Transition: item.Transition}
	}
	var records = make([]config.Record, len(in.GetRecords()))
	for i, item := range in.GetRecords() {
		records[i] = config.Record{ID: item.Id, Name: item.Name, Prefix: item.Prefix}
	}
	var btnDefs = make([]config.BtnDef, len(in.GetBtnDefs()))
	for i, item := range in.GetBtnDefs() {
		btnDefs[i] = config.BtnDef{ID: item.Id, Name: item.Name, BtnName: item.BtnName}
	}

	var cfg = config.ConfigItem{
		UUID:     in.GetUUID(),
		Funnel:   stages,
		Campaign: records,
		BtnTime:  btnDefs,
	}
	if err := srv.config.Update(ctx, srv.storage, cfg, in.GetUpdateFlag()); err != nil {
		return &configSrv.UpdateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &configSrv.UpdateResponse{
		StatusCode: http.StatusOK,
		Msg:        "Configs Updated",
	}, nil
}
