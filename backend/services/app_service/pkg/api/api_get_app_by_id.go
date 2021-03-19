package api

import (
	"context"
	"fmt"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
	userSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/user_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (srv AppService) GetByID(ctx context.Context, request *appSrv.GetByIDRequest) (*appSrv.GetByIDResponse, error) {
	// add tracingID from request to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.GetByID] received get app by id request\n", ctx_value.GetString(ctx, "tracingID"))

	app, err := srv.app.GetByID(ctx, srv.mongoC, request)
	if err != nil {
		return &appSrv.GetByIDResponse{
			StatusCode: http.StatusInternalServerError,
			Msg:        "could not get application data",
			// rest of the fields will take default values
		}, nil
	}

	// collect member data of app from user-service
	// convert []interface of member to []string -> but this is not nice looks for different solution pls
	userUuids := make([]string, len(app["member"].(primitive.A)))
	for i, v := range app["member"].(primitive.A) {
		userUuids[i] = fmt.Sprint(v)
	}
	respUser, err := srv.userService.GetUsersByID(ctx, &userSrv.GetUsersByIDRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		UserUuids:  userUuids,
	})
	if err != nil {
		return &appSrv.GetByIDResponse{
			StatusCode: http.StatusInternalServerError,
			Msg:        "could not get application data",
			// rest of the fields will take default values
		}, nil
	}
	logrus.Infof("User grpc: %v", respUser)
	// I am unhappy because the memory layout of the FullUser and User message should be the same but due to the pointer and slices are not..
	// hence I need to "type cast" them with a loop #un-cool #thisCanBeDoneBetter
	var memberList []*appSrv.FullUser = make([]*appSrv.FullUser, len(respUser.GetUsers()))
	for i, item := range respUser.GetUsers() {
		logrus.Info(item)
		memberList[i] = &appSrv.FullUser{
			Uuid:          item.Uuid,
			Username:      item.Username,
			FirstName:     item.FirstName,
			LastName:      item.LastName,
			OrgnDomain:    item.OrgnDomain,
			OrgnPosition:  item.OrgnPosition,
			ProfileImgUrl: item.ProfileImgUrl,
		}
	}
	// this also is super uncool (all dude to the fact that I don't get how to import common types in freaking protobuf..)
	var settings []string = make([]string, len(app["setting"].(primitive.A)))
	for i, item := range app["setting"].(primitive.A) {
		settings[i] = fmt.Sprint(item)
	}
	return &appSrv.GetByIDResponse{
		StatusCode:  http.StatusOK,
		Msg:         "application data",
		Name:        app["appName"].(string),
		Description: app["appDescription"].(string),
		OwnerUuid:   app["ownerUUID"].(string),
		Member:      memberList,
		Settings:    settings,
	}, nil
}
