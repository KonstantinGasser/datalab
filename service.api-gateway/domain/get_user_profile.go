package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	userproto "github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
)

func (svc gatewaylogic) GetUserProfile(ctx context.Context, uuid string) (*common.UserInfo, errors.ErrApi) {

	resp, err := svc.userClient.Get(ctx, &userproto.GetRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		CallerUuid: uuid,
		ForUuid:    uuid,
	})
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not load User profile",
			Err:    err,
		}
	}
	return resp.GetUser(), nil
}
