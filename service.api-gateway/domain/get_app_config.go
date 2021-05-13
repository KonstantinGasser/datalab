package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	appconfig "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
)

func (svc gatewaylogic) GetAppConfig(ctx context.Context, uuid string) (*common.AppConfigInfo, errors.ErrApi) {
	resp, err := svc.appconfigClient.Get(ctx, &appconfig.GetRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		ForUuid:    uuid,
	})
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "could not get App-Config",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("could not get app config: %s", resp.GetMsg()),
		}
	}
	return resp.GetConfigs(), nil
}
