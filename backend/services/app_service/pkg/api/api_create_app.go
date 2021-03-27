package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/clickstream/backend/protobuf/app_service"
	"github.com/KonstantinGasser/clickstream/backend/services/app_service/pkg/app"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/KonstantinGasser/clickstream/utils/unique"
	"github.com/sirupsen/logrus"
)

// CreateApp is the api endpoint if a user wants to create new app in the system
func (srv AppService) CreateApp(ctx context.Context, request *appSrv.CreateAppRequest) (*appSrv.CreateAppResponse, error) {
	// add tracingID from request to context
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
	logrus.Infof("<%v>[appService.CreateApp] received request\n", ctx_value.GetString(ctx, "tracingID"))

	appUUID, err := unique.UUID()
	if err != nil {
		return &appSrv.CreateAppResponse{StatusCode: http.StatusInternalServerError, Msg: "could not create app", AppUuid: ""}, nil
	}
	status, err := srv.app.CreateApp(ctx, srv.storage, app.AppItem{
		UUID:        appUUID,
		AppName:     request.GetName(),
		Description: request.GetDescription(),
		OwnerUUID:   request.GetOwnerUuid(),
		OrgnDomain:  request.GetOrganization(),
		Member:      append(request.GetMember(), request.GetOwnerUuid()), // by default owner will always be in the app team
		Settings:    request.GetSettings(),
	})
	if err != nil {
		logrus.Errorf("<%v>[appService.CreateApp] could not create app: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.CreateAppResponse{StatusCode: int32(status), Msg: "could not create app"}, nil
	}
	return &appSrv.CreateAppResponse{
		StatusCode: int32(status),
		Msg:        "app has been created",
		AppUuid:    appUUID,
	}, nil
}
