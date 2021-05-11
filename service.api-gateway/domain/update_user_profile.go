package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	userproto "github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/required"
)

type UserUpdateForm struct {
	FirstName     string `json:"first_name" required:"yes"`
	LastName      string `json:"last_name" required:"yes"`
	OrgnPosition  string `json:"orgn_position" required:"yes"`
	ProfileImgURL string `json:"profile_img_url"`
}

func (svc gatewaylogic) UpdateUserProfile(ctx context.Context, uuid string, form UserUpdateForm) errors.ErrApi {
	if err := required.Atomic(&form); err != nil {
		return errors.ErrAPI{
			Status: http.StatusBadRequest,
			Msg:    "All fields are required",
			Err:    err,
		}
	}
	resp, err := svc.userClient.Update(ctx, &userproto.UpdateRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		CallerUuid: uuid,
		User: &userproto.UpdatableUser{
			FirstName:     form.FirstName,
			LastName:      form.LastName,
			OrgnPosition:  form.OrgnPosition,
			ProfileImgUrl: form.ProfileImgURL,
		},
	})
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not update User profile",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("could not update user profile"),
		}
	}
	return nil
}
