package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
func (handler Handler) Create(ctx context.Context, in *proto.CreateRequest) (*proto.CreateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-administer.Create] received request\n", ctx_value.GetString(ctx, "tracingID"))

	appUuid, err := handler.domain.Create(ctx, in)
	if err != nil {
		logrus.Errorf("<%v>[service.app-administer.Create] could not create app: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &proto.CreateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			AppUuid:    "",
		}, nil
	}
	return &proto.CreateResponse{
		StatusCode: http.StatusOK,
		Msg:        "App has been created",
		AppUuid:    appUuid,
	}, nil
}
