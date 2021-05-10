package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Create(ctx context.Context, in *proto.CreateRequest) (*proto.CreateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.user-administer.Create] received request\n", ctx_value.GetString(ctx, "tracingID"))

	err := handler.domain.CreateUser(ctx, in)
	if err != nil {
		return &proto.CreateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.CreateResponse{
		StatusCode: http.StatusOK,
		Msg:        "User-Account created",
	}, nil
}
