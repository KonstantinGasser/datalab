package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/hasher"
	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
)

func (svc gatewaylogic) HasAppPermissions(ctx context.Context, userUuid, appUuid, appName, organization string) errors.ErrApi {
	appHash := hasher.Build(appName, organization)
	resp, err := svc.appClient.MayAcquireToken(ctx, &appsvc.MayAcquireTokenRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		CallerUuid: userUuid,
		AppUuid:    appUuid,
		AppHash:    appHash,
	})
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not check if allowed",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK || !resp.GetIsAllowed() {
		return errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("could not authorize request: %s", resp.GetMsg()),
		}
	}
	return nil
}
