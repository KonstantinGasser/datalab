package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppMetaServer) LockApp(ctx context.Context, in *proto.LockAppRequest) (*proto.LockAppResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.LockApp] received request\n", tracingId)

	err := server.updateService.LockApp(ctx, in.GetAppUuid(), in.GetAuthedUser())
	if err != nil {
		logrus.Errorf("[%v][server.LockApp] could not lock app: %v\n", tracingId, err.Error())
		return &proto.LockAppResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.LockAppResponse{
		StatusCode: http.StatusOK,
		Msg:        "App locked",
	}, nil
}
