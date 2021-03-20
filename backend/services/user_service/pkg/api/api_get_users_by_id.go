package api

import (
	"context"

	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv UserService) GetUsersByID(ctx context.Context, request *userSrv.GetUsersByIDRequest) (*userSrv.GetUsersByIDResponse, error) {
	// add tracingID to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())

	logrus.Infof("<%v>[userService.GetUsersByID] received get users by id request\n", ctx_value.GetString(ctx, "tracingID"))

	status, userList, err := srv.user.GetByIDs(ctx, srv.storage, request.GetUserUuids())
	if err != nil {
		logrus.Errorf("<%v>[userService.GetUsersByID] could not execute GetByIDs: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &userSrv.GetUsersByIDResponse{StatusCode: int32(status), Msg: "Could not get users information", Users: []*userSrv.User{}}, nil
	}

	// convert found userList to grpc User slice
	var users []*userSrv.User = make([]*userSrv.User, len(userList))
	for i, item := range userList {
		users[i] = &userSrv.User{
			Uuid:          item.UUID,
			Username:      item.Username,
			FirstName:     item.FirstName,
			LastName:      item.LastName,
			OrgnDomain:    item.OrgnDomain,
			OrgnPosition:  item.OrgnPosition,
			ProfileImgUrl: item.ProfileImgURL,
		}
	}
	return &userSrv.GetUsersByIDResponse{
		StatusCode: int32(status),
		Msg:        "users record by uuids",
		Users:      users,
	}, nil
}
