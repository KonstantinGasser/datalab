package api

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) AppendMember(ctx context.Context, request *appSrv.AppendMemberRequest) (*appSrv.AppendMemberResponse, error) {
	// add tracingID from request to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.GetApps] received get apps request\n", ctx_value.GetString(ctx, "tracingID"))

	status, err := srv.app.AppendMember(ctx, srv.mongoC, request)
	if err != nil {
		logrus.Errorf("<%v>[appService.AppendMember] could not append member: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.AppendMemberResponse{
			StatusCode: int32(status),
			Msg:        "could not add members to app",
		}, nil
	}
	return &appSrv.AppendMemberResponse{
		StatusCode: int32(status),
		Msg:        "member appended to app",
	}, nil
}
