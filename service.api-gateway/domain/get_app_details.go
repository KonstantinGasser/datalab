package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
)

func (svc gatewaylogic) GetAppDetails(ctx context.Context, userUuid, appUuid string) (*common.AppInfo, errors.ErrApi) {
	resp, err := svc.appClient.Get(ctx, &appsvc.GetRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppUuid:    appUuid,
		CallerUuid: userUuid,
	})
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not load App-Info",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("could not get app info"),
		}
	}
	return resp.GetApp(), nil
}
