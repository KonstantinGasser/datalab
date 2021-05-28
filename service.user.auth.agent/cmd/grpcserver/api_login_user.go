package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server UserAuthServer) Login(ctx context.Context, in *proto.LoginRequest) (*proto.LoginResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Login] received request\n", tracingId)

	accessToken, err := server.authService.Login(ctx,
		in.GetUsername(),
		in.GetPassword(),
	)
	if err != nil {
		logrus.Errorf("[%v][server.Login] could not login user: %v\n", tracingId, err.Error())
		return &proto.LoginResponse{
			StatusCode:  err.Code(),
			Msg:         err.Info(),
			AccessToken: "",
		}, nil
	}
	return &proto.LoginResponse{
		StatusCode:  http.StatusOK,
		Msg:         "User logged in",
		AccessToken: accessToken,
	}, nil
}
