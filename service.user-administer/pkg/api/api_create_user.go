package api

import (
	"context"
	"net/http"
	"strings"

	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/service_user/pkg/user"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/utils/hash"
	"github.com/KonstantinGasser/datalab/utils/unique"
	"github.com/KonstantinGasser/required"
	"github.com/sirupsen/logrus"
)

// CreateUser receives the grpc request and handles user registration
func (srv UserService) Create(ctx context.Context, in *userSrv.CreateRequest) (*userSrv.CreateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[userService.Create] received request\n", ctx_value.GetString(ctx, "tracingID"))

	// create unique user UUID (also used as pk _id for mongo document) and hash password
	uuid, err := unique.UUID()
	if err != nil {
		logrus.Errorf("<%v>[userService.Create] could not generate UUID for user: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.CreateResponse{StatusCode: http.StatusInternalServerError, Msg: "could not create user"}, nil
	}
	hashedPassword, err := hash.FromPassword([]byte(in.GetUser().GetPassword()))
	if err != nil {
		logrus.Errorf("<%v>[userService.Create] could not hash user password: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.CreateResponse{StatusCode: http.StatusInternalServerError, Msg: "could not create user"}, nil
	}

	user := user.UserItem{
		UUID:          uuid,
		Username:      strings.TrimSpace(in.GetUser().GetUsername()),
		Password:      hashedPassword,
		FirstName:     strings.TrimSpace(in.GetUser().GetFirstName()),
		LastName:      strings.TrimSpace(in.GetUser().GetLastName()),
		OrgnDomain:    strings.TrimSpace(in.GetUser().GetOrgnDomain()),
		OrgnPosition:  strings.TrimSpace(in.GetUser().GetOrgnPosition()),
		ProfileImgURL: "https://avatars.githubusercontent.com/u/43576797?v=4", // can be set to default image later
	}
	if err := required.Atomic(&user); err != nil {
		logrus.Errorf("<%v>[userService.Create] new user: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.CreateResponse{StatusCode: http.StatusBadRequest, Msg: "missing mandatory fields"}, nil
	}
	if err := srv.user.Create(ctx, srv.storage, user); err != nil {
		logrus.Errorf("<%v>[userService.Create] could not create user: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &userSrv.CreateResponse{StatusCode: err.Code(), Msg: err.Info()}, nil
	}
	return &userSrv.CreateResponse{StatusCode: http.StatusOK, Msg: "user has been created"}, nil
}
