package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppConfigServer) UnlockConfig(ctx context.Context, in *proto.UnlockConfigRequest) (*proto.UnlockConfigResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.UnlockConfig] received request\n", tracingId)

	err := server.modifyService.UnlockConfig(ctx, in.GetAppRefUuid(), in.GetAuthedUser())
	if err != nil {
		logrus.Errorf("[%v][server.UnlockConfig] could not unlock app config: %v\n", tracingId, err)
		return &proto.UnlockConfigResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.UnlockConfigResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Config Unlocked",
	}, nil
}
