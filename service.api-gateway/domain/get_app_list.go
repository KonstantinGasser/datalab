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

func (svc gatewaylogic) GetAppList(ctx context.Context, uuid string) ([]*common.AppMetaInfo, errors.ErrApi) {

	resp, err := svc.appClient.GetList(ctx, &appsvc.GetListRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		UserClaims: ctx_value.GetAuthedUser(ctx),
	})
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not get App List",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("%v", resp.GetMsg()),
		}
	}
	return resp.GetAppList(), nil
}
