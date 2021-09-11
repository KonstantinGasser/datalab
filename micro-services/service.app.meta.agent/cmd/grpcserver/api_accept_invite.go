package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (s AppMetaServer) AcceptInvite(ctx context.Context, in *proto.AcceptInviteRequest) (*proto.AcceptInviteResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.AcceptInvite] received request\n", tracingId)

	err := s.inviteService.AcceptInvite(ctx,
		in.GetAppUuid(),
	)
	if err != nil {
		logrus.Errorf("[%v][server.AcceptInvite] could not accept invite: %v\n", tracingId, err.Error())
		return &proto.AcceptInviteResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.AcceptInviteResponse{
		StatusCode: http.StatusOK,
		Msg:        "Accepted Invite",
	}, nil
}
