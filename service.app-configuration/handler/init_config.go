package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Init(ctx context.Context, in *proto.InitRequest) (*proto.InitResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-configuration.Init] received request\n", ctx_value.GetString(ctx, "tracingID"))

	err := handler.domain.InitConfigs(ctx, in)
	if err != nil {
		logrus.Errorf("<%v>[service.app-configuration.Init] could not init configs: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &proto.InitResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.InitResponse{
		StatusCode: http.StatusOK,
		Msg:        "App-Configs initialized",
	}, nil
}
