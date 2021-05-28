package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppConfigServer) Initialize(ctx context.Context, in *proto.InitRequest) (*proto.InitResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Initialize] received request\n", tracingId)

	err := server.initService.InitializeAppConfig(ctx, in.GetAppRefUuid(), in.GetAppOwner())
	if err != nil {
		logrus.Errorf("[%v][server.Initialize] could not init App Config: %v\n", tracingId, err.Error())
		return &proto.InitResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.InitResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Config initialized",
	}, nil
}
