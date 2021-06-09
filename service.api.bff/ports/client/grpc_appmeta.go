package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	grpcAppMeta "github.com/KonstantinGasser/datalab/service.app.meta.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

type ClientAppMeta struct {
	Conn grpcAppMeta.AppMetaClient
}

func NewClientAppMeta(clientAddr string) (*ClientAppMeta, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := grpcAppMeta.NewAppMetaClient(conn)
	return &ClientAppMeta{
		Conn: client,
	}, nil
}

func (client ClientAppMeta) CreateApp(ctx context.Context, r *apps.CreateAppRequest) (string, errors.Api) {

	resp, err := client.Conn.Create(ctx, &grpcAppMeta.CreateRequest{
		Tracing_ID:   ctx_value.GetString(ctx, "tracingID"),
		OwnerUuid:    r.OwnerUuid,
		Name:         r.AppName,
		Organization: r.Organization,
		Description:  r.AppDesc,
		AppUrl:       r.AppUrl,
	})
	if err != nil {
		return "", errors.New(http.StatusInternalServerError,
			err,
			"Could not create App")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return "", errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}

	return resp.GetAppUuid(), nil
}

func (client ClientAppMeta) GetApp(ctx context.Context, r *apps.GetAppRequest) (*common.AppInfo, errors.Api) {

	resp, err := client.Conn.Get(ctx, &grpcAppMeta.GetRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AuthedUser: r.AuthedUser,
		AppUuid:    r.AppUuid,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError,
			err,
			"Could not get App")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return resp.GetApp(), nil
}

func (client ClientAppMeta) GetAppList(ctx context.Context, r *apps.GetAppListRequest) ([]*common.AppSubset, errors.Api) {

	resp, err := client.Conn.GetList(ctx, &grpcAppMeta.GetListRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AuthedUser: r.AuthedUser,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError,
			err,
			"Could not get Apps")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return resp.GetAppList(), nil
}

func (client ClientAppMeta) SendInvite(ctx context.Context, r *apps.SendInviteRequest) (string, errors.Api) {

	resp, err := client.Conn.Invite(ctx, &grpcAppMeta.InviteRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AuthedUser: r.AuthedUser,
		AppUuid:    r.AppUuid,
		UserUuid:   r.InvitedUuid,
	})
	if err != nil {
		return "", errors.New(http.StatusInternalServerError,
			err,
			"Could not send App Invite")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return "", errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return resp.GetAppName(), nil
}

func (client ClientAppMeta) InviteReminderOK(ctx context.Context, r *apps.InviteReminderRequest) errors.Api {

	resp, err := client.Conn.InviteReminderOK(ctx, &grpcAppMeta.InviteReminderOKRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppUuid:    r.AppUuid,
		UserUuid:   r.UserUuid,
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not check if Reminder can be send")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return nil
}

func (client ClientAppMeta) AcceptInvite(ctx context.Context, r *apps.AcceptInviteRequest) errors.Api {

	resp, err := client.Conn.AcceptInvite(ctx, &grpcAppMeta.AcceptInviteRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AuthedUser: r.AuthedUser,
		AppUuid:    r.AppUuid,
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not accept App Invite")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return nil
}

func (client ClientAppMeta) LockApp(ctx context.Context, appUuid string, authedUser *common.AuthedUser) errors.Api {
	resp, err := client.Conn.LockApp(ctx, &grpcAppMeta.LockAppRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AuthedUser: authedUser,
		AppUuid:    appUuid,
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not lock app")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return nil
}

func (client ClientAppMeta) UnlockApp(ctx context.Context, r *apps.UnlockRequest) error {
	resp, err := client.Conn.UnlockApp(ctx, &grpcAppMeta.UnlockAppRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AuthedUser: r.AuthedUser,
		AppUuid:    r.AppUuid,
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not unlock app")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return nil
}
