package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) IsAuthed(ctx context.Context, in *proto.IsAuthedRequest) (*proto.IsAuthedResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.user-authentication.IsAuthed] received request\n", ctx_value.GetString(ctx, "tracingID"))

	claims, err := handler.domain.IsAuthenticated(ctx, in)
	if err != nil {
		return &proto.IsAuthedResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			IsAuthed:   false,
			Claims:     nil,
		}, nil
	}
	return &proto.IsAuthedResponse{
		StatusCode: http.StatusOK,
		Msg:        "User is authenticated",
		IsAuthed:   true,
		Claims:     claims,
	}, nil
}
