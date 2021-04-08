package api

import (
	"context"
	"net/http"

	appSrv "github.com/KonstantinGasser/datalabs/backend/protobuf/app_service"
	"github.com/KonstantinGasser/datalabs/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

// GetApps serves as the grpc implementation to retrieve all apps created by the logged in user
<<<<<<< HEAD
func (srv AppService) GetAppList(ctx context.Context, in *appSrv.GetAppListRequest) (*appSrv.GetAppListResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", in.GetTracing_ID())
=======
func (srv AppService) GetList(ctx context.Context, request *appSrv.GetListRequest) (*appSrv.GetListResponse, error) {
	ctx = ctx_value.AddValue(ctx, "tracingID", request.GetTracing_ID())
>>>>>>> feature_app_token
	logrus.Infof("<%v>[appService.GetAppList] received request\n", ctx_value.GetString(ctx, "tracingID"))

	status, result, err := srv.app.GetAppList(ctx, srv.storage, in.GetCallerUuid())
	if err != nil {
		logrus.Errorf("<%v>[appService.GetAppList] could not get apps: %v\n", ctx_value.GetString(ctx, "tracingID"), err)
		return &appSrv.GetListResponse{StatusCode: int32(status), Msg: "could not get list of apps", AppList: []*appSrv.SimpleApp{}}, nil
	}

	// convert appList to grpc LightApp slice
	var apps []*appSrv.SimpleApp = make([]*appSrv.SimpleApp, len(result))
	for i, item := range result {
		apps[i] = &appSrv.SimpleApp{Name: item.AppName, Uuid: item.UUID}
	}
	return &appSrv.GetListResponse{
		StatusCode: http.StatusOK,
		Msg:        "apps found",
		AppList:    apps,
	}, nil
}
