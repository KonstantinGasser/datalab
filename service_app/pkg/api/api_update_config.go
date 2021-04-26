package api

import (
	"context"

	appSrv "github.com/KonstantinGasser/datalabs/protobuf/app_service"
	"github.com/KonstantinGasser/datalabs/service_app/pkg/config"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) UpdateCfg(ctx context.Context, in *appSrv.UpdateCfgRequest) (*appSrv.UpdateCfgResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.UpdateCfg] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, ok, err := srv.app.HasPermissions(ctx, srv.storage, in.GetCallerUuid(), in.GetAppUuid())
	if err != nil || !ok {
		logrus.Warn("<%v>[appService.UpdateCfg] could not authorize request: %v - %v\n", ctx_value.GetString(ctx, "tracingID"), ok, err)
		return &appSrv.UpdateCfgResponse{StatusCode: int32(status), Msg: "could not process request"}, nil
	}

	var funnel = []config.Stage{}
	for _, item := range in.GetFunnel() {
		funnel = append(funnel, config.Stage{
			ID:         item.Id,
			Name:       item.Name,
			Transition: item.Transition,
		})
	}
	var campaign = []config.Record{}
	for _, item := range in.GetCampaign() {
		campaign = append(campaign, config.Record{
			ID:     item.Id,
			Name:   item.Name,
			Prefix: item.Prefix,
		})
	}
	var btnTime = []config.BtnDef{}
	for _, item := range in.GetBtnTime() {
		btnTime = append(btnTime, config.BtnDef{
			ID:      item.Id,
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

	status, err = srv.config.UpdateByFlag(ctx, srv.storage, cfg, in.GetUpdateFlag())
	if err != nil {
		logrus.Errorf("<%v>[appService.UpdateCfg] could not update config(s)\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.UpdateCfgResponse{StatusCode: int32(status), Msg: "could not update configs"}, nil
	}
	return &appSrv.UpdateCfgResponse{StatusCode: int32(status), Msg: "configs updated"}, nil
}
