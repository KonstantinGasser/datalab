package api

import (
	"context"
	"net/http"
	"strings"

	userSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/user_service"
	"github.com/KonstantinGasser/datalabs/backend/services/user_service/pkg/user"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/KonstantinGasser/datalabs/utils/hash"
	"github.com/KonstantinGasser/datalabs/utils/unique"
	"github.com/sirupsen/logrus"
)

// CreateUser receives the grpc request and handles user registration
func (srv UserService) Create(ctx context.Context, in *userSrv.CreateRequest) (*userSrv.CreateResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
	logrus.Infof("<%v>[userService.CreateUser] received request\n", ctx_value.GetString(ctx, "tracingID"))

	// create unique user UUID (also used as pk _id for mongo document) and hash password
	uuid, err := unique.UUID()
	if err != nil {
		logrus.Errorf("<%v>[userService.CreateUser] could not generate UUID for user: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.CreateResponse{StatusCode: http.StatusInternalServerError, Msg: "could not create user"}, nil
	}
	hashedPassword, err := hash.FromPassword([]byte(in.GetUser().GetPassword()))
	if err != nil {
		logrus.Errorf("<%v>[userService.CreateUser] could not hash user password: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.CreateResponse{StatusCode: http.StatusInternalServerError, Msg: "could not create user"}, nil
	}

	status, err := srv.user.InsertNew(ctx, srv.storage, user.UserItem{
		UUID:          uuid,
		Username:      strings.TrimSpace(in.GetUser().GetUsername()),
		Password:      hashedPassword,
		FirstName:     strings.TrimSpace(in.GetUser().GetFirstName()),
		LastName:      strings.TrimSpace(in.GetUser().GetLastName()),
		OrgnDomain:    strings.TrimSpace(in.GetUser().GetOrgnDomain()),
		OrgnPosition:  strings.TrimSpace(in.GetUser().GetOrgnPosition()),
		ProfileImgURL: "https://avatars.githubusercontent.com/u/43576797?v=4", // can be set to default image later
	})
	if err != nil {
		logrus.Errorf("<%v>[userService.CreateUser] could not create user: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.CreateResponse{StatusCode: int32(status), Msg: "could not create user"}, nil
	}
	return &userSrv.CreateResponse{StatusCode: int32(status), Msg: "user has been created"}, nil
}
