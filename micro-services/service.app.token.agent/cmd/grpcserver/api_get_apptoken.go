package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppTokenServer) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Get] received request\n", tracingId)

	appToken, err := server.fetchService.FetchById(ctx, in.GetAppUuid())
	if err != nil {
		logrus.Errorf("[%v][server.Get] could not get app token: %v\n", tracingId, err.Error())
		return &proto.GetResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Token:      nil,
		}, nil
	}
	return &proto.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "Fetched App Token",
		Token: &common.AppAccessToken{
			Locked:     appToken.Locked,
			Jwt:        appToken.Jwt,
			Expiration: appToken.Exp,
		},
	}, nil
}
