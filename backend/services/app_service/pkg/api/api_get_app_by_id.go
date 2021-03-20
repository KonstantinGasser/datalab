package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (srv AppService) GetByID(ctx context.Context, request *appSrv.GetByIDRequest) (*appSrv.GetByIDResponse, error) {
	// add tracingID from request to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.GetByID] received get app by id request\n", ctx_value.GetString(ctx, "tracingID"))

	status, app, err := srv.app.GetByID(ctx, srv.storage, request.GetAppUuid())
	if err != nil {
		logrus.Errorf("<%v>[appService.GetByID] could not GetByID: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.GetByIDResponse{StatusCode: int32(status), Msg: "could not get application data"}, nil
	}
	// ask user-service for some more details about each member of the app
	respUser, err := srv.userService.GetUsersByID(ctx, &userSrv.GetUsersByIDRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		UserUuids:  app.Member,
	})
	if err != nil || respUser.GetStatusCode() >= 400 {
		return &appSrv.GetByIDResponse{
			StatusCode:  http.StatusOK,
			Msg:         "application data",
			Name:        app.AppName,
			Description: app.Description,
			OwnerUuid:   app.OwnerUUID,
			Member:      nil,
			Settings:    app.Settings,
		}, nil
	}
	// I am unhappy because the memory layout of the FullUser and User message should be the same but due to the pointer and slices are not..
	// hence I need to "type cast" them with a loop #un-cool #thisCanBeDoneBetter
	var memberList []*appSrv.FullUser = make([]*appSrv.FullUser, len(respUser.GetUsers()))
	for i, item := range respUser.GetUsers() {
		memberList[i] = &appSrv.FullUser{
			Uuid:          item.GetUuid(),
			Username:      item.GetUsername(),
			FirstName:     item.GetFirstName(),
			LastName:      item.GetLastName(),
			OrgnDomain:    item.GetOrgnDomain(),
			OrgnPosition:  item.GetOrgnPosition(),
			ProfileImgUrl: item.GetProfileImgUrl(),
		}
	}
	return &appSrv.GetByIDResponse{
		StatusCode:  http.StatusOK,
		Msg:         "application data",
		Name:        app.AppName,
		Description: app.Description,
		OwnerUuid:   app.OwnerUUID,
		Member:      memberList,
		Settings:    app.Settings,
	}, nil
}
