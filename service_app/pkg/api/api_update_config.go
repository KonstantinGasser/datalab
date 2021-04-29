package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	"github.com/KonstantinGasser/datalab/service_app/pkg/config"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) UpdateCfg(ctx context.Context, in *appSrv.UpdateCfgRequest) (*appSrv.UpdateCfgResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.UpdateCfg] received request\n", ctx_value.GetString(ctx, "tracingID"))

	if err := srv.app.HasPermissions(ctx, srv.storage, in.GetCallerUuid(), in.GetAppUuid()); err != nil {
		return &appSrv.UpdateCfgResponse{StatusCode: err.Code(), Msg: err.Info()}, nil
	}
	var funnel = []config.Stage{}
	for _, item := range in.GetStages() {
		funnel = append(funnel, config.Stage{
			ID:         item.GetId(),
			Name:       item.Name,
			Transition: item.Transition,
		})
	}
	var campaign = []config.Record{}
	for _, item := range in.GetRecords() {
		campaign = append(campaign, config.Record{
			ID:     item.GetId(),
			Name:   item.Name,
			Prefix: item.Prefix,
		})
	}
	var btnTime = []config.BtnDef{}
	for _, item := range in.GetBtnDefs() {
		btnTime = append(btnTime, config.BtnDef{
			ID:      item.GetId(),
			Name:    item.Name,
			BtnName: item.BtnName,
		})
	}

	cfg := config.Cfgs{
		RefApp:   in.GetAppUuid(),
		Funnel:   funnel,
		Campaign: campaign,
		BtnTime:  btnTime,
	}

	if err := srv.app.UpdateConfig(ctx, srv.storage, cfg, in.GetUpdateFlag()); err != nil {
		logrus.Errorf("<%v>[appService.UpdateCfg] could not update config(s)\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &appSrv.UpdateCfgResponse{StatusCode: err.Code(), Msg: err.Info()}, nil
	}
	return &appSrv.UpdateCfgResponse{StatusCode: http.StatusOK, Msg: "configs updated"}, nil
}
