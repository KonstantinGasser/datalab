package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Invite(ctx context.Context, in *proto.InviteRequest) (*proto.InviteResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-administer.Invite] received request\n", ctx_value.GetString(ctx, "tracingID"))

	appName, appOwner, err := handler.domain.InviteToApp(ctx, in)
	if err != nil {
		return &proto.InviteResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.InviteResponse{
		StatusCode: http.StatusOK,
		Msg:        "Invitation saved",
		AppName:    appName,
		OwnerUuid:  appOwner,
	}, nil
}
