package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Update(ctx context.Context, in *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-configuration.Update] received request\n", ctx_value.GetString(ctx, "tracingID"))

	err := handler.domain.UpdateConfig(ctx, in)
	if err != nil {
		return &proto.UpdateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.UpdateResponse{
		StatusCode: http.StatusOK,
		Msg:        "App-Config has been updated",
	}, nil
}
