package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	userproto "github.com/KonstantinGasser/datalab/service.user-administer/proto"
	authproto "github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/required"
)

type RegisterForm struct {
	Username     string `json:"username" required:"yes"`
	FirstName    string `json:"first_name" required:"yes"`
	LastName     string `json:"last_name" required:"yes"`
	Password     string `json:"password" required:"yes"`
	OrgnDomain   string `json:"orgn_domain" required:"yes"`
	OrgnPosition string `json:"orgn_position" required:"yes"`
}

func (svc gatewaylogic) RegisterUser(ctx context.Context, form RegisterForm) errors.ErrApi {

	if err := required.Atomic(&form); err != nil {
		return errors.ErrAPI{
			Status: http.StatusBadRequest,
			Msg:    "All fields are mandatory",
			Err:    err,
		}
	}
	// request user-auth to create and generate user
	authResp, err := svc.userauthClient.Register(ctx, &authproto.RegisterRequest{
		Tracing_ID:   ctx_value.GetString(ctx, "tracingID"),
		Username:     form.Username,
		Password:     form.Password,
		Organisation: form.OrgnDomain,
	})
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create User-Account",
			Err:    err,
		}
	}
	if authResp.GetStatusCode() != http.StatusOK {
		return errors.ErrAPI{
			Status: authResp.GetStatusCode(),
			Msg:    authResp.GetMsg(),
			Err:    fmt.Errorf("user-auth service failed to register new user account"),
		}
	}
	// request user-administer service to create new user account
	userUuid := authResp.GetUserUuid()
	userResp, err := svc.userClient.Create(ctx, &userproto.CreateRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		User: &common.UserInfo{
			Uuid:         userUuid,
			Username:     form.Username,
			FirstName:    form.FirstName,
			LastName:     form.LastName,
			OrgnDomain:   form.OrgnDomain,
			OrgnPosition: form.OrgnPosition,
		},
	})
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create User-Account",
			Err:    err,
		}
	}
	if userResp.GetStatusCode() != http.StatusOK {
		return errors.ErrAPI{
			Status: userResp.GetStatusCode(),
			Msg:    userResp.GetMsg(),
			Err:    fmt.Errorf("user-admin service failed to create new user account"),
		}
	}
	return nil
}
