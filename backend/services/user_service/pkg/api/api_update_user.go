package api

import (
	"context"
	"strings"

	userSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/backend/services/user_service/pkg/user"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// AuthUser is a public interface of the service allowing to authenticate
// a user by its credentials
<<<<<<< HEAD
func (srv UserService) UpdateUser(ctx context.Context, in *userSrv.UpdateUserRequest) (*userSrv.UpdateUserResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
=======
func (srv UserService) Update(ctx context.Context, request *userSrv.UpdateRequest) (*userSrv.UpdateResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
>>>>>>> feature_app_token
	logrus.Infof("<%v>[userService.UpdateUser] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, err := srv.user.Update(ctx, srv.storage, user.UserItemUpdateable{
		UUID:          in.GetCallerUuid(),
		FirstName:     strings.TrimSpace(in.GetUser().GetFirstName()),
		LastName:      strings.TrimSpace(in.GetUser().GetLastName()),
		OrgnPosition:  strings.TrimSpace(in.GetUser().GetOrgnPosition()),
		ProfileImgURL: strings.TrimSpace(in.GetUser().GetProfileImgUrl()),
	})
	if err != nil {
		logrus.Errorf("<%v>[userService.UpdateUser] could not update user: %v\n", err)
		return &userSrv.UpdateResponse{StatusCode: int32(status), Msg: "could not update user details"}, nil
	}
	return &userSrv.UpdateResponse{StatusCode: int32(status), Msg: "user details updated"}, nil
}
