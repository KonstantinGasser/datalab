package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	cfgsvc "github.com/KonstantinGasser/datalab/service.app-configuration/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
)

type UpdateConfigForm struct {
	Uuid    string             `json:"app_uuid" required:"yes"`
	Stages  []*common.Funnel   `json:"stages"`
	Records []*common.Campaign `json:"records"`
	BtnDefs []*common.BtnTime  `json:"btn_defs"`
}

func (svc gatewaylogic) UpdateAppConfig(ctx context.Context, form UpdateConfigForm, flag string) errors.ErrApi {

	resp, err := svc.appconfigClient.Update(ctx, &cfgsvc.UpdateRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		UUID:       form.Uuid,
		UpdateFlag: flag,
		Stages:     form.Stages,
		Records:    form.Records,
		BtnDefs:    form.BtnDefs,
	})
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not update App Configs",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("%s", resp.GetMsg()),
		}
	}
	return nil
}
