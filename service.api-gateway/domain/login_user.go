package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	authproto "github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/required"
)

type LoginForm struct {
	Username string `json:"username" required:"yes"`
	Password string `json:"password" required:"yes"`
}

func (svc gatewaylogic) LoginUser(ctx context.Context, form LoginForm) (string, errors.ErrApi) {

	if err := required.Atomic(&form); err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusBadRequest,
			Msg:    "All fields are mandatory",
			Err:    err,
		}
	}

	loginResp, err := svc.userauthClient.Login(ctx, &authproto.LoginRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		Username:   form.Username,
		Password:   form.Password,
	})
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not login User",
			Err:    err,
		}
	}
	if loginResp.GetStatusCode() != http.StatusOK {
		return "", errors.ErrAPI{
			Status: loginResp.GetStatusCode(),
			Msg:    loginResp.GetMsg(),
			Err:    fmt.Errorf("could not login user"),
		}
	}
	return loginResp.GetJwt(), nil
}
