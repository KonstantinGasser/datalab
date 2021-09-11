package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	authproto "github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
)

func (svc notificationlogic) IsLoggedIn(ctx context.Context, token string) (*common.AuthedUser, errors.Api) {

	resp, err := svc.userauthClient.IsAuthed(ctx, &authproto.IsAuthedRequest{
		Tracing_ID:  ctx_value.GetString(ctx, "tracingID"),
		AccessToken: token,
	})
	if err != nil {
		return nil, errors.API{
			Status: http.StatusInternalServerError,
			Msg:    "Could not authenticate request",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK || !resp.GetIsAuthed() {
		return nil, errors.API{
			Status: http.StatusUnauthorized,
			Msg:    "User is not authenticated",
			Err:    err,
		}
	}
	return resp.GetAuthedUser(), nil
}
