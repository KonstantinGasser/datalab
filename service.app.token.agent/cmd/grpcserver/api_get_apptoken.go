package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppTokenServer) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	tracingId := ctx.Value("tracingID")
	logrus.Infof("[%v][server.Get] received request\n", tracingId)

	authedUser := in.GetAuthedUser()
	if authedUser == nil {
		return &proto.GetResponse{
			StatusCode: http.StatusUnauthorized,
			Msg:        "User is not authroized to get App Token",
			Token:      nil,
		}, nil
	}
	appToken, err := server.fetchService.FetchById(
		ctx, in.GetAppUuid(),
		authedUser.Uuid,
		authedUser.GetReadWriteApps()...,
	)
	if err != nil {
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
			Jwt:        appToken.Jwt,
			Expiration: appToken.Exp.Unix(),
		},
	}, nil
}
