package domain

import (
	"context"
	"fmt"
	"net/http"

	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	userauthsvc "github.com/KonstantinGasser/datalab/service.user-authentication/proto"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

type AcceptInviteForm struct {
	AppUuid        string `json:"app_uuid"`
	EventTimestamp int64  `json:"event_timestamp"`
}

func (svc gatewaylogic) AcceptInvite(ctx context.Context, form AcceptInviteForm) (string, errors.ErrApi) {

	resp, err := svc.appClient.AcceptInvite(ctx, &appsvc.AcceptInviteRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		UserClaims: ctx_value.GetAuthedUser(ctx),
		AppUuid:    form.AppUuid,
	})
	fmt.Printf("Resp: %+v\n", resp)
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not process accepting of invitation",
			Err:    err,
		}
	}
	if resp.GetStatusCode() != http.StatusOK {
		return "", errors.ErrAPI{
			Status: resp.GetStatusCode(),
			Msg:    resp.GetMsg(),
			Err:    fmt.Errorf("%s", resp.GetMsg()),
		}
	}
	fmt.Println("before add access")
	// needs to move to Kafka ASAP
	// once the invite status is changed to accept the users permissions
	// need to be appended accordingly
	permissionResp, err := svc.userauthClient.AddAppAccess(ctx, &userauthsvc.AddAppAccessRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		UserUuid:   ctx_value.GetAuthedUser(ctx).GetUuid(),
		AppUuid:    form.AppUuid,
		AppRole:    1, // the role concept has been drop as of now (once on the app you can do all) - might change in later version
	})
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not add app permissions",
			Err:    err,
		}
	}
	if permissionResp.GetStatusCode() != http.StatusOK {
		return "", errors.ErrAPI{
			Status: permissionResp.GetStatusCode(),
			Msg:    permissionResp.GetMsg(),
			Err:    fmt.Errorf("%s", permissionResp.GetMsg()),
		}
	}
	fmt.Println("after add access")
	// needs to move to Kafka ASAP
	err = svc.RemoveNotification(ctx, RemoveEvent{
		UserUuid: ctx_value.GetAuthedUser(ctx).GetUuid(),
		Timesamp: form.EventTimestamp,
	})
	if err != nil {
		logrus.Errorf("remove notification: %v\n", err)
	}
	return permissionResp.GetUpdatedToken(), nil
}
