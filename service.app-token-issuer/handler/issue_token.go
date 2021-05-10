package handler

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) Issue(ctx context.Context, in *proto.IssueRequest) (*proto.IssueResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[service.app-token-issuer.Issue] received request\n", ctx_value.GetString(ctx, "tracingID"))

	token, err := handler.domain.IssueToken(ctx, in)
	if err != nil {
		logrus.Errorf("<%v>[service.app-token-issuer.Issue] could not issue app-token: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &proto.IssueResponse{
			StatusCode: err.Code(),
			Msg:        err.Info(),
			Token:      nil,
		}, nil
	}
	return &proto.IssueResponse{
		StatusCode: http.StatusOK,
		Msg:        "App-Token has been issued",
		Token:      token,
	}, nil
}
