package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) VerifySameOrgn(ctx context.Context, in *userSrv.VerifySameOrgnRequest) (*userSrv.VerifySameOrgnResposne, error) {
	// ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	// logrus.Infof("<%v>[userService.VerifySameOrgn] received request\n", ctx_value.GetString(ctx, "tracingID"))

	// status, areValid, missedItems, err := srv.user.CompareOrgn(ctx, srv.storage, in.GetBaseObject(), in.GetCompareWith())
	// if err != nil {
	// 	logrus.Errorf("<%v>[userService.VerifySameOrgn] could not compare items: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
	// 	return &userSrv.VerifySameOrgnResposne{StatusCode: int32(status), Msg: "could not compare items"}, nil
	// }
	// return &userSrv.VerifySameOrgnResposne{
	// 	StatusCode:    int32(status),
	// 	Msg:           "all items have been compared",
	// 	TruthfulValid: areValid,
	// 	InvalidList:   missedItems,
	// }, nil
	return nil, nil
}
