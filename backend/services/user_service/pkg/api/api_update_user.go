package api

import (
	"context"
	"net/http"
	"strings"

	userSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/backend/services/user_service/pkg/user"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
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
	if err := required.All(&updatableUser); err != nil {
		logrus.Errorf("<%v>[userService.Update] new user: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.UpdateResponse{StatusCode: http.StatusBadRequest, Msg: "missing of mandatory field"}, nil
	}
	status, err := srv.user.Update(ctx, srv.storage, updatableUser)
	if err != nil {
		logrus.Errorf("<%v>[userService.Update] could not update user: %v\n", err)
		return &userSrv.UpdateResponse{StatusCode: int32(status), Msg: "could not update user details"}, nil
	}
	return &userSrv.UpdateResponse{StatusCode: int32(status), Msg: "user details updated"}, nil
}
