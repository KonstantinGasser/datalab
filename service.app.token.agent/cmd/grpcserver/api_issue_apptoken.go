package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppTokenServer) Issue(ctx context.Context, in *proto.IssueRequest) (*proto.IssueResponse, error) {
	tracingId := ctx.Value("tracingID")
	logrus.Infof("[%v][server.Issue] received request\n", tracingId)

	jwt, exp, err := server.modifySevice.IssueAppToken(ctx, in.GetAppUuid(), in.GetCallerUuid())
	if err != nil {
		return &proto.IssueResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Token:      nil,
		}, nil
	}
	return &proto.IssueResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Token updated",
		Token: &common.AppAccessToken{
			Jwt:        jwt,
			Expiration: exp.Unix(),
		},
	}, nil
}
