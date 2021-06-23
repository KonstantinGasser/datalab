package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server UserMetaServer) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Get] received request\n", tracingId)

	user, err := server.fetchService.FetchLoggedIn(ctx)
	if err != nil {
		logrus.Errorf("[%v][server.Get] could not get user: %v\n", tracingId, err.Error())
		return &proto.GetResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			User:       nil,
		}, nil
	}
	return &proto.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "User profile",
		User: &common.UserInfo{
			Uuid:         user.Uuid,
			Username:     user.Username,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			OrgnDomain:   user.Organization,
			OrgnPosition: user.Position,
			Avatar:       user.Avatar,
		},
	}, nil
}

func (server UserMetaServer) GetById(ctx context.Context, in *proto.GetByIdRequest) (*proto.GetByIdResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.GetById] received request\n", tracingId)

	user, err := server.fetchService.FetchById(ctx, in.GetUserUuid())
	if err != nil {
		logrus.Errorf("[%v][server.Get] could not get user: %v\n", tracingId, err.Error())
		return &proto.GetByIdResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			User:       nil,
		}, nil
	}
	return &proto.GetByIdResponse{
		StatusCode: http.StatusOK,
		Msg:        "User profile",
		User: &common.UserInfo{
			Uuid:         user.Uuid,
			Username:     user.Username,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			OrgnDomain:   user.Organization,
			OrgnPosition: user.Position,
			Avatar:       user.Avatar,
		},
	}, nil
}
