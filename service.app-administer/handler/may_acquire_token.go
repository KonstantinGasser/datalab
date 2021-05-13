package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
func (handler Handler) MayAcquireToken(ctx context.Context, in *proto.MayAcquireTokenRequest) (*proto.MayAcquireTokenResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-administer.MayAcquireToken] received request\n", ctx_value.GetString(ctx, "tracingID"))

	if err := handler.domain.MayAcquireToken(ctx, in); err != nil {
		logrus.Infof("<%v>[service.app-administer.MayAcquireToken] could not verify if ok: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &proto.MayAcquireTokenResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			IsAllowed:  false,
		}, nil
	}
	return &proto.MayAcquireTokenResponse{
		StatusCode: http.StatusOK,
		Msg:        "caller is allowed to acquire app-token",
		IsAllowed:  true,
	}, nil
}
