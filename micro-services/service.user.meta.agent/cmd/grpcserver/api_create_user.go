package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server UserMetaServer) Create(ctx context.Context, in *proto.CreateRequest) (*proto.CreateResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Create] received request\n", tracingId)

	err := server.createService.CreateUser(ctx,
		in.GetUser().GetUuid(),
		in.GetUser().GetUsername(),
		in.GetUser().GetFirstName(),
		in.GetUser().GetLastName(),
		in.GetUser().GetOrgnDomain(),
		in.GetUser().GetOrgnPosition(),
	)
	if err != nil {
		logrus.Errorf("[%v][server.Create] could not create user: %v\n", tracingId, err.Error())
		return &proto.CreateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.CreateResponse{
		StatusCode: http.StatusOK,
		Msg:        "User Profile created",
	}, nil
}
