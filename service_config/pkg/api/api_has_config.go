package api

import (
	"context"
	"net/http"

	configSrv "github.com/KonstantinGasser/datalab/protobuf/config_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv ConfigServer) HasConfigs(ctx context.Context, in *configSrv.HasConfigsRequest) (*configSrv.HasConfigsResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[configService.HasConfigs] received request\n", ctx_value.GetString(ctx, "tracingID"))

	ok, err := srv.config.HasAtLeastOne(ctx, srv.storage, in.GetUUID())
	if err != nil {
		logrus.Infof("<%v>[configService.HasConfigs] could not check if configs set: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &configSrv.HasConfigsResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Ok:         false,
		}, nil
	}
	if !ok {
		return &configSrv.HasConfigsResponse{
			StatusCode: http.StatusNotFound,
			Msg:        "No configs found",
			Ok:         false,
		}, nil
	}
	return &configSrv.HasConfigsResponse{
		StatusCode: http.StatusOK,
		Msg:        "At least one config set",
		Ok:         true,
	}, nil
}
