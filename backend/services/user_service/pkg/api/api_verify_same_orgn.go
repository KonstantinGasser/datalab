package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) VerifySameOrgn(ctx context.Context, request *userSrv.VerifySameOrgnRequest) (*userSrv.VerifySameOrgnResposne, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[userService.VerifySameOrgn] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, areValid, missedItems, err := srv.user.VerifySameOrgn(ctx, srv.storage, request.GetBaseObject(), request.GetCompareWith())
	if err != nil {
		logrus.Errorf("<%v>[userService.VerifySameOrgn] could not compare items: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.VerifySameOrgnResposne{StatusCode: int32(status), Msg: "could not compare items"}, nil
	}
	return &userSrv.VerifySameOrgnResposne{
		StatusCode:    int32(status),
		Msg:           "all items have been compared",
		TruthfulValid: areValid,
		InvalidList:   missedItems,
	}, nil
}
