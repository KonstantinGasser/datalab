package api

import (
	"context"
	"net/http"

	userSrv "github.com/KonstantinGasser/datalab/protobuf/user_service"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv UserService) Get(ctx context.Context, in *userSrv.GetRequest) (*userSrv.GetResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())

	logrus.Infof("<%v>[userService.GetUser] received request\n", ctx_value.GetString(ctx, "tracingID"))

	user, err := srv.user.Get(ctx, srv.storage, in.GetForUuid())
	if err != nil {
		logrus.Errorf("<%v>[userService.GetUser] could not get user details: %v\n", ctx_value.GetString(ctx, "tracingID"), err.Error())
		return &userSrv.GetResponse{StatusCode: err.Code(), Msg: err.Info(), User: nil}, nil
	}
	return &userSrv.GetResponse{
		StatusCode: http.StatusOK,
		Msg:        "requested user found",
		User: &userSrv.ComplexUser{
			Uuid:          user.UUID,
			Username:      user.Username,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			OrgnDomain:    user.OrgnDomain,
			OrgnPosition:  user.OrgnPosition,
			ProfileImgUrl: user.ProfileImgURL,
		},
	}, nil
}
