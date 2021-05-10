package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	apptokenSrv "github.com/KonstantinGasser/datalab/protobuf/apptoken_service"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/apptoken"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppTokenServer) Issue(ctx context.Context, in *apptokenSrv.IssueRequest) (*apptokenSrv.IssueResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.Issue] received request\n", ctx_value.GetString(ctx, "tracingID"))

	respIsAllowed, err := srv.app.IsAllowedToGenToken(ctx, &appSrv.IsAllowedToGenTokenRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		CallerUuid: in.GetCallerUuid(),
		AppUuid:    in.GetAppUuid(),
		AppHash:    in.GetAppHash(),
	})
	if err != nil {
		logrus.Errorf("<%v>[appService.Issue] could not execute grpc.IsAllowedToGenToken: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &apptokenSrv.IssueResponse{
			StatusCode: http.StatusInternalServerError,
			Msg:        "Could not generate App-Token",
		}, nil
	}
	if respIsAllowed.GetStatusCode() != http.StatusOK || !respIsAllowed.GetIsAllowed() {
		logrus.Errorf("<%v>[appService.Issue] request is not authorized to generate app token\n", ctx_value.GetString(ctx, "tracingID"))
		return &apptokenSrv.IssueResponse{
			StatusCode: http.StatusUnauthorized,
			Msg:        "No permissions to create app token",
			Token:      nil,
		}, nil
	}

	token, tokenErr := srv.apptoken.Issue(ctx, srv.storage, apptoken.TokenClaims{
		AppUuid:   in.GetAppUuid(),
		AppHash:   in.GetAppHash(),
		AppOrigin: in.GetAppOrigin(),
	})
	if tokenErr != nil {
		logrus.Errorf("<%v>[appService.Issue] could not issue apptoken: %v\n", ctx_value.GetString(ctx, "tracingID"), tokenErr.Error())
		return &apptokenSrv.IssueResponse{
			StatusCode: tokenErr.Code(),
			Msg:        tokenErr.Info(),
			Token:      nil,
		}, nil
	}
	return &apptokenSrv.IssueResponse{
		StatusCode: http.StatusOK,
		Msg:        "Created App-Token",
		Token: &apptokenSrv.MetaToken{
			Token: token,
			Exp:   12,
		},
	}, nil
}
