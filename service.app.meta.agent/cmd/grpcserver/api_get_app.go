package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppMetaServer) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Get] received request\n", tracingId)

	app, err := server.fechtService.FetchAppByID(ctx,
		in.GetAppUuid(),
		// in.GetAuthedUser(),
	)
	if err != nil {
		logrus.Errorf("[%v][server.Get] could not get App: %v\n", tracingId, err.Error())
		return &proto.GetResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			App:        nil,
		}, nil
	}

	// convert from App.Member to protobuf common.AppMember
	var members = make([]*common.AppMember, len(app.Members))
	for i, member := range app.Members {
		members[i] = &common.AppMember{
			Uuid:   member.Uuid,
			Status: int32(member.Status),
		}
	}
	return &proto.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "App found",
		App: &common.AppInfo{
			Uuid:        app.Uuid,
			Locked:      app.Locked,
			Name:        app.Name,
			URL:         app.URL,
			Description: app.Description,
			Owner:       app.OwnerUuid,
			Member:      members,
		},
	}, nil
}

func (server AppMetaServer) GetList(ctx context.Context, in *proto.GetListRequest) (*proto.GetListResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.GetList] received request\n", tracingId)

	apps, err := server.fechtService.FetchAppSubsets(ctx)
	if err != nil {
		logrus.Errorf("[%v][server.GetList] could not get Apps: %v\n", tracingId, err.Error())
		return &proto.GetListResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			AppList:    nil,
		}, nil
	}

	// convert AppSubset to protobuf common.AppSubset
	var appSubsets = make([]*common.AppSubset, len(apps))
	for i, app := range apps {
		appSubsets[i] = &common.AppSubset{
			Uuid: app.Uuid,
			Name: app.Name,
		}
	}
	return &proto.GetListResponse{
		StatusCode: http.StatusOK,
		Msg:        "Apps related to User",
		AppList:    appSubsets,
	}, nil
}
