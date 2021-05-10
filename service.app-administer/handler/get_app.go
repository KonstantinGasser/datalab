package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
func (handler Handler) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-administer.Get] received request\n", ctx_value.GetString(ctx, "tracingID"))

	app, cfg, token, err := handler.domain.GetSingle(ctx, in)
	if err != nil {
		return &proto.GetResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "Collected App Information",
		App:        app,
		AppConfig:  cfg,
		AppToken:   token,
	}, nil
}
