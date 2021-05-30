package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server UserMetaServer) Update(ctx context.Context, in *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Upate] received request\n", tracingId)

	err := server.updateService.UpadeUser(ctx,
		in.GetCallerUuid(),
		in.GetUser().GetFirstName(),
		in.GetUser().GetLastName(),
		in.GetUser().GetOrgnPosition(),
	)
	if err != nil {
		logrus.Errorf("[%v][server.Update] could not update user: %v\n", tracingId, err.Error())
		return &proto.UpdateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.UpdateResponse{
		StatusCode: http.StatusOK,
		Msg:        "User profile updated",
	}, nil
}
