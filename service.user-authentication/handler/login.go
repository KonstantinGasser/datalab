package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Login(ctx context.Context, in *proto.LoginRequest) (*proto.LoginResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.user-authentication.Login] received request\n", ctx_value.GetString(ctx, "tracingID"))

	token, err := handler.domain.LoginUser(ctx, in)
	if err != nil {
		logrus.Infof("<%v>[service.user-authentication.Login] could not authenticate user: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &proto.LoginResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Jwt:        "",
		}, nil
	}
	return &proto.LoginResponse{
		StatusCode: http.StatusOK,
		Msg:        "User authenticated",
		Jwt:        token,
	}, nil
}
