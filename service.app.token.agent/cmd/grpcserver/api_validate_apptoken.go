package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppTokenServer) Validate(ctx context.Context, in *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Validate] received request\n", tracingId)

	appUuid, appOrigin, err := server.validateService.ValidateAppToken(ctx,
		in.GetAppToken(),
	)
	if err != nil {
		logrus.Errorf("[%v][server.Validate] could not validate App Token: %v\n", tracingId, err.Error())
		return &proto.ValidateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.ValidateResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Token authorized",
		AppUuid:    appUuid,
		AppOrigin:  appOrigin,
	}, nil
}
