package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) GetColleagues(ctx context.Context, in *proto.GetColleaguesRequest) (*proto.GetColleaguesResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.user-administer.GetColleagues] received request\n", ctx_value.GetString(ctx, "tracingID"))

	users, err := handler.domain.GetColleagues(ctx, in)
	if err != nil {
		return &proto.GetColleaguesResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Colleagues: nil,
		}, nil
	}
	return &proto.GetColleaguesResponse{
		StatusCode: http.StatusOK,
		Msg:        "Colleagues found",
		Colleagues: users,
	}, nil
}
