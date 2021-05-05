package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	apptokenSrv "github.com/KonstantinGasser/datalab/protobuf/apptoken_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppTokenServer) Issue(ctx context.Context, in *apptokenSrv.IssueRequest) (*apptokenSrv.IssueResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.CreateApp] received request\n", ctx_value.GetString(ctx, "tracingID"))

	respIsAllowed, err := srv.app.IsAllowedToGenToken(ctx, &appSrv.IsAllowedToGenTokenRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		CallerUuid: in.GetCallerUuid(),
		AppUuid:    in.GetAppUuid(),
	})
	if err != nil {
		logrus.Errorf("<%v>[appService.CreateApp] could not execute grpc.IsAllowedToGenToken: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &apptokenSrv.IssueResponse{
			StatusCode: http.StatusInternalServerError,
			Msg:        "Could not generate App-Token",
		}, nil
	}
	if respIsAllowed.GetStatusCode() != http.StatusOK {
		return &apptokenSrv.IssueResponse{
			StatusCode: http.StatusUnauthorized,
			Msg:        "No permissions to create app token",
			Token:      "",
		}, nil
	}
	return nil, nil
}
