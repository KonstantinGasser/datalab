package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppTokenServer) Initialize(ctx context.Context, in *proto.InitializeRequest) (*proto.InitializeResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Initialize] received request\n", tracingId)

	err := server.initService.InitializeAppToken(ctx,
		in.GetAppRefUuid(),
		in.GetAppHash(),
		in.GetAppOwner(),
		in.GetAppOrigin(),
	)
	if err != nil {
		logrus.Errorf("[%v][server.Initialize] could not init app token: %v\n", err.Error(), tracingId)
		return &proto.InitializeResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.InitializeResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Token initialized",
	}, nil
}
