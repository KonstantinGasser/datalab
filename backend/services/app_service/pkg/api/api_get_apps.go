package api

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
)

func (srv AppService) GetApps(ctx context.Context, request *appSrv.GetAppsRequest) (*appSrv.GetAppsResponse, error) {
	return nil, nil
}
