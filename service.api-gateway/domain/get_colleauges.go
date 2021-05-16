package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	usersvc "github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
)

func (svc gatewaylogic) GetColleagues(ctx context.Context, userUuid string) ([]*common.UserInfo, errors.ErrApi) {

	resp, err := svc.userClient.GetColleagues(ctx, &usersvc.GetColleaguesRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		UserUuid:   userUuid,
	})
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not load colleagues",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("%s", resp.GetMsg()),
		}
	}
	return resp.GetColleagues(), nil
}
