package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user-permissions/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Init(ctx context.Context, in *proto.InitRequest) (*proto.InitResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.user-permissions.Init] received request\n", ctx_value.GetString(ctx, "tracingID"))

	err := handler.domain.InitPermissions(ctx, in)
	if err != nil {
		logrus.Errorf("<%v>[service.user-permissions.Init] could not init user-permissions\n", ctx_value.GetString(ctx, "tracingID"))
		return &proto.InitResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.InitResponse{
		StatusCode: http.StatusOK,
		Msg:        "User Permissions initialized",
	}, nil
}
