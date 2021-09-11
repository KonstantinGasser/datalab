package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server UserAuthServer) IsAuthed(ctx context.Context, in *proto.IsAuthedRequest) (*proto.IsAuthedResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.IsAuthed] received request\n", tracingId)

	authedUser, err := server.authService.Authenticate(ctx, in.GetAccessToken())
	if err != nil {
		logrus.Errorf("[%v][server.IsAuthed] could not autenticate user: %v\n", tracingId, err.Error())
		return &proto.IsAuthedResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			IsAuthed:   false,
			AuthedUser: nil,
		}, nil
	}
	permisions, err := server.fetchService.GetById(ctx, authedUser.Uuid)
	if err != nil {
		logrus.Errorf("[%v][server.IsAuthed] could not get permissions of user: %v\n", tracingId, err.Error())
		return &proto.IsAuthedResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			IsAuthed:   false,
			AuthedUser: nil,
		}, nil
	}
	return &proto.IsAuthedResponse{
		StatusCode: http.StatusOK,
		Msg:        "User logged in",
		IsAuthed:   true,
		AuthedUser: &common.AuthedUser{
			Uuid:          authedUser.Uuid,
			Username:      authedUser.Username,
			Organization:  authedUser.Organization,
			ReadWriteApps: permisions.AllowedApps(),
		},
	}, nil
}
