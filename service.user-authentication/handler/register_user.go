package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// Register is the gRPC Entry-point to register (create) a new user record
func (handler Handler) Register(ctx context.Context, in *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.user-authentication.Register] received request\n", ctx_value.GetString(ctx, "tracingID"))

	uuid, err := handler.domain.RegisterNewUser(ctx, in)
	if err != nil {
		logrus.Infof("<%v>[service.user-authentication.Register] could not register user account: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &proto.RegisterResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			UserUuid:   "",
		}, nil
	}
	return &proto.RegisterResponse{
		StatusCode: http.StatusOK,
		Msg:        "User-Account created",
		UserUuid:   uuid,
	}, nil
}
