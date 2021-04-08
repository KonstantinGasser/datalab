package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/app_service"
	userSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) AddMember(ctx context.Context, request *appSrv.AddMemberRequest) (*appSrv.AddMemberResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.AddMember] received request\n", ctx_value.GetString(ctx, "tracingID"))

	// verify that all passed members belong to the app_owner organization
	resp, err := srv.userService.VerifySameOrgn(ctx, &userSrv.VerifySameOrgnRequest{
		Tracing_ID:  ctx_value.GetString(ctx, "tracingID"),
		CallerUuid:  request.GetCallerUuid(),
		BaseObject:  request.GetCallerUuid(),
		CompareWith: request.GetMember(),
	})
	if err != nil {
		logrus.Errorf("<%v>[appService.AddMember] could not execute grpc.AreInSameOrng: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.AddMemberResponse{StatusCode: int32(http.StatusInternalServerError), Msg: "could not verify members orgn"}, nil
	}
	// validate response from user service
	// if not allValid => members which are about to be added to app have different organization which is not allowed
	if !resp.GetTruthfulValid() {
		logrus.Errorf("<%v>[appService.AddMember] not all member share same orgn with owner\nInvalid: %v\n", ctx_value.GetString(ctx, "tracingID"), resp.GetInvalidList())
		return &appSrv.AddMemberResponse{StatusCode: int32(http.StatusBadRequest), Msg: "all member must be from owner orgn"}, nil
	}

	status, err := srv.app.AddMember(ctx, srv.storage, request.GetCallerUuid(), request.GetAppUuid(), request.GetMember())
	if err != nil {
		logrus.Errorf("<v>[appService.AddMember] could not add member to app: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.AddMemberResponse{StatusCode: int32(status), Msg: "could not add member(s) to app"}, nil
	}
	return &appSrv.AddMemberResponse{
		StatusCode: int32(status),
		Msg:        "added member(s) to app",
	}, nil
}
