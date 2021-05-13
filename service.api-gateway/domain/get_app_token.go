package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	apptokissuer "github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
)

func (svc gatewaylogic) GetAppToken(ctx context.Context, uuid string) (*common.AppTokenInfo, errors.ErrApi) {

	resp, err := svc.apptokenClient.Get(ctx, &apptokissuer.GetRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppUuid:    uuid,
	})
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "could not get App-Token",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("could not get app token: %s", resp.GetMsg()),
		}
	}
	return resp.GetToken(), nil
}
