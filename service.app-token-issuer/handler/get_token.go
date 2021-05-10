package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-token-issuer.Get] received request\n", ctx_value.GetString(ctx, "tracingID"))

	token, err := handler.domain.GetToken(ctx, in)
	if err != nil {
		return &proto.GetResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Token:      nil,
		}, nil
	}
	return &proto.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "Found App-Token",
		Token:      token,
	}, nil
}
