package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server UserAuthServer) Register(ctx context.Context, in *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Register] received request\n", tracingId)

	userUuid, err := server.authService.Register(ctx,
		in.GetUsername(),
		in.GetOrganisation(),
		in.GetPassword(),
	)
	if err != nil {
		logrus.Errorf("[%v][server.Register] could not register user: %v\n", tracingId, err.Error())
		return &proto.RegisterResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			UserUuid:   "",
		}, nil
	}
	return &proto.RegisterResponse{
		StatusCode: http.StatusOK,
		Msg:        "User Account created",
		UserUuid:   userUuid,
	}, nil
}
