package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (s AppMetaServer) InviteReminderOK(ctx context.Context, in *proto.InviteReminderOKRequest) (*proto.InviteReminderOKResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.InviteReminderOK] received request\n", tracingId)

	err := s.inviteService.SendInviteReminderOK(ctx, in.GetAppUuid(), in.GetUserUuid())
	if err != nil {
		logrus.Errorf("[%v][server.InviteReminderOK] could not check if send reminder ok: %v\n", tracingId, err.Error())
		return &proto.InviteReminderOKResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.InviteReminderOKResponse{
		StatusCode: http.StatusOK,
		Msg:        "Invite can be re-send",
	}, nil
}
