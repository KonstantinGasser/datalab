package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Update(ctx context.Context, in *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.user-administer.Update] received request\n", ctx_value.GetString(ctx, "tracingID"))

	err := handler.domain.UpdateUser(ctx, in)
	if err != nil {
		logrus.Errorf("<%v>[service.user-administer.Update] could not update user: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &proto.UpdateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.UpdateResponse{
		StatusCode: http.StatusOK,
		Msg:        "User has been updated",
	}, nil
}
