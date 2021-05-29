package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.user-administer.Get] received request\n", ctx_value.GetString(ctx, "tracingID"))

	user, err := handler.domain.GetUser(ctx, in)
	if err != nil {
		return &proto.GetResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			User:       nil,
		}, nil
	}
	return &proto.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "Found user",
		User:       user,
	}, nil
}
