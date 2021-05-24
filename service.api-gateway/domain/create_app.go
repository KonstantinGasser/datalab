package domain

import (
	"context"
	"fmt"
	"net/http"

	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/required"
)

type CreateAppForm struct {
	Name        string `json:"app_name" required:"yes"`
	URL         string `json:"app_URL" required:"yes"`
	Description string `json:"app_description" required:"yes"`
}

func (svc gatewaylogic) CreateApp(ctx context.Context, uuid, organization string, form CreateAppForm) (string, errors.ErrApi) {
	if err := required.Atomic(&form); err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusBadGateway,
			Msg:    "All fields required",
			Err:    err,
		}
	}
	resp, err := svc.appClient.Create(ctx, &appsvc.CreateRequest{
		Tracing_ID:   ctx_value.GetString(ctx, "tracingID"),
		OwnerUuid:    uuid,
		Name:         form.Name,
		Organization: organization,
		Description:  form.Description,
		AppUrl:       form.URL,
	})
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create App",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return "", errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("could not create app"),
		}
	}
	return resp.GetAppUuid(), nil
}
