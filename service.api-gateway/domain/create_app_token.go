package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	apptokensvc "github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
)

func (svc gatewaylogic) CreateAppToken(ctx context.Context, userUuid, appUuid, appOrigin, appHash string) (*common.AppTokenInfo, errors.ErrApi) {

	resp, err := svc.apptokenClient.Issue(ctx, &apptokensvc.IssueRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		CallerUuid: userUuid,
		AppUuid:    appUuid,
		AppHash:    appHash,
		AppOrigin:  appOrigin,
	})
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create App Token",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("could not create app token: %s", resp.GetMsg()),
		}
	}
	return resp.GetToken(), nil
}
