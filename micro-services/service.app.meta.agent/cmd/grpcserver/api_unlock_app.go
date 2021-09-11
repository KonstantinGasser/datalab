package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppMetaServer) UnlockApp(ctx context.Context, in *proto.UnlockAppRequest) (*proto.UnlockAppResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.UnlockApp] received request\n", tracingId)

	err := server.updateService.UnlockApp(ctx, in.GetAppUuid())
	if err != nil {
		logrus.Errorf("[%v][server.UnlockApp] could not unlock app: %v\n", tracingId, err.Error())
		return &proto.UnlockAppResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.UnlockAppResponse{
		StatusCode: http.StatusOK,
		Msg:        "App locked",
	}, nil
}
