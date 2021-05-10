package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-configuration.Get] received request\n", ctx_value.GetString(ctx, "tracingID"))

	cfgs, err := handler.domain.GetConfigs(ctx, in)
	if err != nil {
		logrus.Errorf("<%v>[service.app-configuration.Get] could not get configs: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &proto.GetResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Configs:    nil,
		}, nil
	}
	return &proto.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "Found App-Configs",
		Configs:    cfgs,
	}, nil
}
