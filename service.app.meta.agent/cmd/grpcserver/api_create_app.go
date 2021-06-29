package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppMetaServer) Create(ctx context.Context, in *proto.CreateRequest) (*proto.CreateResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Create] received request\n", tracingId)

	appUuid, err := server.createSerivce.CreateDefaultApp(ctx,
		in.GetName(),
		in.GetAppUrl(),
		in.GetOwnerUuid(),
		in.GetOrganization(),
		in.GetDescription(),
		in.GetIsPrivate(),
	)
	if err != nil {
		logrus.Errorf("[%v][server.Create] could not create app: %v\n", tracingId, err.Error())
		return &proto.CreateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			AppUuid:    "",
		}, nil
	}
	return &proto.CreateResponse{
		StatusCode: http.StatusOK,
		Msg:        "App created",
		AppUuid:    appUuid,
	}, nil
}
