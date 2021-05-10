package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
func (handler Handler) Delete(ctx context.Context, in *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-administer.Delete] received request\n", ctx_value.GetString(ctx, "tracingID"))

	err := handler.domain.Delete(ctx, in)
	if err != nil {
		logrus.Infof("<%v>[service.app-administer.Delete] could not delete app: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &proto.DeleteResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.DeleteResponse{
		StatusCode: http.StatusOK,
		Msg:        "App has been deleted",
	}, nil
}
