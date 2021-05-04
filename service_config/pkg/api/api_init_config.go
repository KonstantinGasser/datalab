package api

import (
	"context"
	"net/http"

	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv ConfigServer) Init(ctx context.Context, in *configSrv.InitRequest) (*configSrv.InitResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[configService.InitConfigs] received request\n", ctx_value.GetString(ctx, "tracingID"))

	uuid, err := srv.config.Init(ctx, srv.storage)
	if err != nil {
		logrus.Infof("<%v>[configService.InitConfigs] could not init config: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &configSrv.InitResponse{StatusCode: err.Code(), Msg: err.Info(), UUID: ""}, nil
	}
	return &configSrv.InitResponse{
		StatusCode: http.StatusOK,
		Msg:        "Configs init done",
		UUID:       uuid,
	}, nil
}
