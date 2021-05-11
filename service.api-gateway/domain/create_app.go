package domain

import (
	"context"
	"net/http"

	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/required"
)

type CreateAppForm struct {
	Name        string   `json:"app_name" required:"yes"`
	URL         string   `json:"app_URL" required:"yes"`
	Description string   `json:"app_description" required:"yes"`
	Member      []string `json:"app_member"`
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
		Member:       nil,
	})
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create App",
			Err:    err,
		}
	}
	return resp.GetAppUuid(), nil
}
