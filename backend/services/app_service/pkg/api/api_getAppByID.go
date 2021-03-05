package api

import (
	"context"

	appSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/app_service"
)

func (srv AppService) GetByID(ctx context.Context, request *appSrv.GetByIDRequest) (*appSrv.GetByIDResponse, error) {
	return nil, nil
}
