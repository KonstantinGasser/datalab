package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) AddMember(ctx context.Context, request *appSrv.AddMemberRequest) (*appSrv.AddMemberResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.AddMember] received request\n", ctx_value.GetString(ctx, "tracingID"))

	// verify that all passed members belong to the app_owner organization
	// compare_to is the owner of the app and all values(users) must be from the owner.orgn
	resp, err := srv.userService.AreInSameOrgn(ctx, &userSrv.AreInSameOrgnRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		CompareTo:  request.GetOwnerUuid(),
		Values:     request.GetMember(),
	})
	if err != nil {
		logrus.Errorf("<%v>[appService.AddMember] could not execute grpc.AreInSameOrng: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.AddMemberResponse{StatusCode: int32(http.StatusInternalServerError), Msg: "could not verify members orgn"}, nil
	}
	// validate response from user service
	// if not allValid => members which are about to be added to app have different organization which is not allowed
	if !resp.GetTruthfulValid() {
		logrus.Errorf("<%v>[appService.AddMember] not all member share same orgn with owner\nInvalid: %v\n", ctx_value.GetString(ctx, "tracingID"), resp.GetMissMatches())
		return &appSrv.AddMemberResponse{StatusCode: int32(http.StatusBadRequest), Msg: "all member must be from owner orgn"}, nil
	}

	status, err := srv.app.AddMember(ctx, srv.storage, request.GetOwnerUuid(), request.GetAppUuid(), request.GetMember())
	if err != nil {
		logrus.Errorf("<v>[appService.AddMember] could not add member to app: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.AddMemberResponse{StatusCode: int32(status), Msg: "could not add member(s) to app"}, nil
	}
	return &appSrv.AddMemberResponse{
		StatusCode: int32(status),
		Msg:        "added member(s) to app",
	}, nil
}
