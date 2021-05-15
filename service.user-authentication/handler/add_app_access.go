package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) AddAppAccess(ctx context.Context, in *proto.AddAppAccessRequest) (*proto.AddAppAccessResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.user-authentication.AddAppAccess] received request\n", ctx_value.GetString(ctx, "tracingID"))

	err := handler.domain.AddAppAccess(ctx, in)
	if err != nil {
		logrus.Errorf("<%v>[service.user-authentication.AddAppAccess] could not add permission: %v\n",
			ctx_value.GetString(ctx, "tracingID"),
			err.Error())
		return &proto.AddAppAccessResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.AddAppAccessResponse{
		StatusCode: http.StatusOK,
		Msg:        "App Permissions Updated",
	}, nil
}
