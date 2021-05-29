package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (s AppMetaServer) Invite(ctx context.Context, in *proto.InviteRequest) (*proto.InviteResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Invite] received request\n", tracingId)

	appName, err := s.inviteService.SendInvite(ctx,
		in.GetAppUuid(),
		in.GetUserUuid(),
		in.GetAuthedUser(),
	)
	if err != nil {
		return &proto.InviteResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			AppName:    "",
			OwnerUuid:  "",
		}, nil
	}
	return &proto.InviteResponse{
		StatusCode: http.StatusOK,
		Msg:        "Invited User",
		AppName:    appName,
		OwnerUuid:  in.GetAuthedUser().GetUuid(),
	}, nil
}
