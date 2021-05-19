package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
func (handler Handler) GetList(ctx context.Context, in *proto.GetListRequest) (*proto.GetListResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-administer.GetList] received request\n", ctx_value.GetString(ctx, "tracingID"))

	apps, err := handler.domain.GetMultiple(ctx, in)
	if err != nil {
		logrus.Errorf("<%v>[service.app-administer.GetList] could not get list: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &proto.GetListResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			AppList:    nil,
		}, nil
	}
	return &proto.GetListResponse{
		StatusCode: http.StatusOK,
		Msg:        "Found related Apps",
		AppList:    apps,
	}, nil
}
