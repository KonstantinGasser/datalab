package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/sirupsen/logrus"
)

func (server AppConfigServer) Get(ctx context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Get] received request\n", tracingId)

	appConfig, err := server.fetchService.GetById(ctx, in.GetAppUuid())
	if err != nil {
		logrus.Errorf("[%v][server.Get] could not get App Config: %v\n", tracingId, err.Error())
		return &proto.GetResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Configs:    nil,
		}, nil
	}

	// translate AppConfig struct to protobuf struct
	var stages = make([]*common.Stage, len(appConfig.Funnel))
	for i, item := range appConfig.Funnel {
		stages[i] = &common.Stage{Id: item.Id, Name: item.Name, Transition: item.Transition, Regex: item.Regex, Trigger: item.Trigger}
	}
	var records = make([]*common.Record, len(appConfig.Campaign))
	for i, item := range appConfig.Campaign {
		records[i] = &common.Record{Id: item.Id, Name: item.Name, Suffix: item.Suffix}
	}
	var btnDefs = make([]*common.BtnDef, len(appConfig.BtnTime))
	for i, item := range appConfig.BtnTime {
		btnDefs[i] = &common.BtnDef{Id: item.Id, Name: item.Name, BtnName: item.BtnName}
	}
	return &proto.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Config fetched",
		Configs: &common.AppConfigurations{
			Locked:   appConfig.Locked,
			Funnel:   stages,
			Campaign: records,
			BtnTime:  btnDefs,
		},
	}, nil
}

func (server AppConfigServer) GetForClient(ctx context.Context, in *proto.GetForClientRequest) (*proto.GetForClientResponse, error) {
	logrus.Infof("[-][server.GetForClient] received request\n")

	config, err := server.libFetchService.FromStore(ctx, in.GetAppUuid())
	if err != nil {
		return &proto.GetForClientResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Config:     nil,
		}, nil
	}

	var stages = make([]*proto.ClientStage, len(config.Stages))
	for i, item := range config.Stages {
		stages[i] = &proto.ClientStage{Type: int32(item.Type), Transition: item.Transition, Regex: item.Regex}
	}
	var btnDefs = make([]*proto.ClientBtnDef, len(config.BtnDefs))
	for i, item := range config.BtnDefs {
		btnDefs[i] = &proto.ClientBtnDef{Name: item.Name}
	}

	return &proto.GetForClientResponse{
		StatusCode: http.StatusOK,
		Msg:        "Client Configs",
		Config: &proto.ClientConfig{
			Stages:  stages,
			BtnDefs: btnDefs,
		},
	}, nil
}
