package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) GetList(ctx context.Context, in *proto.GetListRequest) (*proto.GetListResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.user-administer.GetList] received request\n", ctx_value.GetString(ctx, "tracingID"))

	users, err := handler.domain.GetUsers(ctx, in)
	if err != nil {
		return &proto.GetListResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			UserList:   nil,
		}, nil
	}
	return &proto.GetListResponse{
		StatusCode: http.StatusOK,
		Msg:        "Found users",
		UserList:   users,
	}, nil
}
