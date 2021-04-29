package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalab/protobuf/app_service"
	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) AddMember(ctx context.Context, in *appSrv.AddMemberRequest) (*appSrv.AddMemberResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[appService.AddMember] received in\n", ctx_value.GetString(ctx, "tracingID"))

	resp, err := srv.userService.VerifySameOrgn(ctx, &userSrv.VerifySameOrgnRequest{
		Tracing_ID:  ctx_value.GetString(ctx, "tracingID"),
		CallerUuid:  in.GetCallerUuid(),
		BaseObject:  in.GetCallerUuid(),
		CompareWith: in.GetMember(),
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

	if err := srv.app.AddMember(ctx, srv.storage, in.GetCallerUuid(), in.GetAppUuid(), in.GetMember()); err != nil {
		logrus.Errorf("<v>[appService.AddMember] %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &appSrv.AddMemberResponse{StatusCode: err.Code(), Msg: "could not add member(s) to app"}, nil
	}
	return &appSrv.AddMemberResponse{
		StatusCode: http.StatusOK,
		Msg:        "added member(s) to app",
	}, nil
}
