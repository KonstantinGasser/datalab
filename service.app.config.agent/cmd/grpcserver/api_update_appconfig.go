package grpcserver

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.config.agent/internal/appconfigs"
	"github.com/sirupsen/logrus"
)

func (server AppConfigServer) Update(ctx context.Context, in *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	tracingId := in.GetTracing_ID()
	logrus.Infof("[%v][server.Update] received request\n", tracingId)

	var err errors.Api
	switch in.GetUpdateFlag() {
	case "funnel":
		var stages = make([]appconfigs.Stage, len(in.GetStages()))
		for i, item := range in.GetStages() {
			stages[i] = appconfigs.Stage{
				Id:         item.Id,
				Name:       item.Name,
				Transition: item.Transition,
				Trigger:    int32(item.Trigger),
			}
		}
		err = server.modifyService.UpdateFunnel(ctx, in.GetAppRefUuid(), stages)
	case "campaign":
		var records = make([]appconfigs.Record, len(in.GetRecords()))
		for i, item := range in.GetRecords() {
			records[i] = appconfigs.Record{
				Id:     item.Id,
				Name:   item.Name,
				Suffix: item.Suffix,
			}
		}
		err = server.modifyService.UpdateCampaign(ctx, in.GetAppRefUuid(), records)
	case "btntime":
		var btnDefs = make([]appconfigs.BtnDef, len(in.GetBtnDefs()))
		for i, item := range in.GetBtnDefs() {
			btnDefs[i] = appconfigs.BtnDef{
				Id:      item.Id,
				Name:    item.Name,
				BtnName: item.BtnName,
			}
		}
		err = server.modifyService.UpdateBtnTime(ctx, in.GetAppRefUuid(), btnDefs)
	default:
		logrus.Errorf("[%v][server.Update] invalid update flag: %s\n", tracingId, in.GetUpdateFlag())
		return &proto.UpdateResponse{
			StatusCode: http.StatusBadRequest,
			Msg:        "Provided update flag is invalid",
		}, nil
	}
	if err != nil {
		logrus.Errorf("[%v][server.Update] could not update App Config: %v\n", tracingId, err.Error())
		return &proto.UpdateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.UpdateResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Config updated",
	}, nil
}
