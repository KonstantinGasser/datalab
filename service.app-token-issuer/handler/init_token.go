package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Init(ctx context.Context, in *proto.InitRequest) (*proto.InitResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-token-issuer.Init] received request\n", ctx_value.GetString(ctx, "tracingID"))

	err := handler.domain.InitToken(ctx, in)
	if err != nil {
		return &proto.InitResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.InitResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Token initialized",
	}, nil
}
