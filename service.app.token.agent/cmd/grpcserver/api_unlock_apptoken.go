package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppTokenServer) UnlockAppToken(ctx context.Context, in *proto.UnlockAppTokenRequest) (*proto.UnlockAppTokenResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.UnlockAppToken] received request\n", tracingId)

	err := server.modifySevice.UnlockAppToken(ctx, in.GetAppUuid(), in.GetAuthedUser())
	if err != nil {
		logrus.Errorf("[%v][server.UnlockAppToken] could not unlock App Token: %v\n", tracingId, err.Error())
		return &proto.UnlockAppTokenResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.UnlockAppTokenResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Token unlocked",
	}, nil
}
