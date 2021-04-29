package api

import (
	"context"
	"net/http"
	"strings"

	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/service_user/pkg/user"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/required"
	"github.com/sirupsen/logrus"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
func (srv UserService) Update(ctx context.Context, in *userSrv.UpdateRequest) (*userSrv.UpdateResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[userService.Update] received request\n", ctx_value.GetString(ctx, "tracingID"))

	updatableUser := user.UserItemUpdateable{
		UUID:          in.GetCallerUuid(),
		FirstName:     strings.TrimSpace(in.GetUser().GetFirstName()),
		LastName:      strings.TrimSpace(in.GetUser().GetLastName()),
		OrgnPosition:  strings.TrimSpace(in.GetUser().GetOrgnPosition()),
		ProfileImgURL: strings.TrimSpace(in.GetUser().GetProfileImgUrl()),
	}
	required.Debug(&updatableUser).Pretty()
	if err := required.Atomic(&updatableUser); err != nil {
		logrus.Errorf("<%v>[userService.Update] new user: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.UpdateResponse{StatusCode: http.StatusBadRequest, Msg: "missing of mandatory field"}, nil
	}
	if err := srv.user.Update(ctx, srv.storage, updatableUser); err != nil {
		logrus.Errorf("<%v>[userService.Update] could not update user: %v\n", err.Error())
		return &userSrv.UpdateResponse{StatusCode: err.Code(), Msg: err.Info()}, nil
	}
	return &userSrv.UpdateResponse{StatusCode: http.StatusOK, Msg: "user details updated"}, nil
}
