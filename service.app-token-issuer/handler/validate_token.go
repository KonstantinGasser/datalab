package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Validate(ctx context.Context, in *proto.ValidateRequest) (*proto.ValidateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-token-issuer.Validate] received request\n", ctx_value.GetString(ctx, "tracingID"))

	appUuid, appOrigin, err := handler.domain.ValidateToken(ctx, in)
	if err != nil {
		logrus.Errorf("<%v>[service.app-token-issuer.Issue] could not issue app-token: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &proto.ValidateResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
		}, nil
	}
	return &proto.ValidateResponse{
		StatusCode: http.StatusOK,
		Msg:        "App-Token authenticated and authorized",
		AppUuid:    appUuid,
		AppOrigin:  appOrigin,
	}, nil
}
