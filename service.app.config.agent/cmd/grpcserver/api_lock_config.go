package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppConfigServer) LockConfig(ctx context.Context, in *proto.LockConfigRequest) (*proto.LockConfigResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.LockConfig] received request\n", tracingId)

	err := server.modifyService.LockConfig(ctx, in.GetAppRefUuid())
	if err != nil {
		logrus.Errorf("[%v][server.LockConfig] could not lock app config: %v\n", tracingId, err)
		return &proto.LockConfigResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.LockConfigResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Config Locked",
	}, nil
}
