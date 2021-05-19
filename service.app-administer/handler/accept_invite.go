package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) AcceptInvite(ctx context.Context, in *proto.AcceptInviteRequest) (*proto.AcceptInviteResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-administer.AcceptInvite] received request\n", ctx_value.GetString(ctx, "tracingID"))

	err := handler.domain.AcceptInvite(ctx, in)
	if err != nil {
		logrus.Infof("<%v>[service.app-administer.AcceptInvite] could not process request: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &proto.AcceptInviteResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.AcceptInviteResponse{
		StatusCode: http.StatusOK,
		Msg:        "Invite Accepted",
	}, nil
}
