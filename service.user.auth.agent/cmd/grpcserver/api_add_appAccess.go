package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server UserAuthServer) AddAppAccess(ctx context.Context, in *proto.AddAppAccessRequest) (*proto.AddAppAccessResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.AddAppAccess] received request\n", tracingId)

	err := server.addService.AddApp(ctx, in.GetUserUuid(), in.GetAppUuid())
	if err != nil {
		logrus.Errorf("[%v][server.AddAppAccess] could not add app access: %v\n", tracingId, err.Error())
		return &proto.AddAppAccessResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.AddAppAccessResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Access added",
	}, nil
}
