package domain

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/KonstantinGasser/required"
)

type InviteForm struct {
	InvitedUuid string `json:"invited_uuid"`
	AppUuid     string `json:"app_uuid"`
	AppRole     int    `json:"app_role"`
}

func (svc gatewaylogic) InviteToAppProcess(ctx context.Context, form InviteForm) errors.ErrApi {
	if err := required.Atomic(&form); err != nil {
		return errors.ErrAPI{
			Status: http.StatusBadRequest,
			Msg:    "Missing field values",
			Err:    err,
		}
	}
	user := ctx_value.GetAuthedUser(ctx)
	if user == nil {
		return errors.ErrAPI{
			Status: http.StatusUnauthorized,
			Msg:    "User is not authenticated",
			Err:    fmt.Errorf("no authed-user found in context"),
		}
	}
	inviteResp, err := svc.appClient.Invite(ctx, &appsvc.InviteRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		UserClaims: ctx_value.GetAuthedUser(ctx),
		AppUuid:    form.AppUuid,
		UserUuid:   form.InvitedUuid,
	})
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not invite user",
			Err:    err,
		}
	}
	if inviteResp.GetStatusCode() != http.StatusOK {
		return errors.ErrAPI{
			Status: inviteResp.GetStatusCode(),
			Msg:    inviteResp.GetMsg(),
			Err:    fmt.Errorf("%s", inviteResp.GetMsg()),
		}
	}

	owner, ownerErr := svc.GetUserProfile(ctx, inviteResp.GetOwnerUuid())
	if ownerErr != nil {
		return ownerErr
	}

	notification := Notification{
		ReceiverUuid: form.InvitedUuid,
		ReceiverOrgn: user.GetOrganization(),
		Mutation:     "APP_INVITE",
		Value: map[string]interface{}{
			"event":     0,
			"app_name":  inviteResp.GetAppName(),
			"app_uuid":  form.AppUuid,
			"app_owner": strings.Join([]string{owner.GetFirstName(), owner.GetLastName()}, " "),
		},
	}
	notificationErr := svc.IssueNotification(ctx, notification)
	if err != nil {
		return notificationErr
	}
	return nil
}
