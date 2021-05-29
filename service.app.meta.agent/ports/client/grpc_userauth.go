package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports"
	grpcUserAuth "github.com/KonstantinGasser/datalab/service.user.auth.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

type ClientUserAuth struct {
	Conn grpcUserAuth.UserAuthenticationClient
}

func NewClientUserAuth(clientAddr string) (*ClientUserAuth, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := grpcUserAuth.NewUserAuthenticationClient(conn)
	return &ClientUserAuth{
		Conn: client,
	}, nil
}

func (client ClientUserAuth) AddAppAccess(ctx context.Context, memberUuid, appUuid string) errors.Api {
	resp, err := client.Conn.AddAppAccess(ctx, &grpcUserAuth.AddAppAccessRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		UserUuid:   memberUuid,
		AppUuid:    appUuid,
		AppRole:    *common.AppRole_EDITOR.Enum(),
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError, err, "Could not add App Permissions")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(), fmt.Errorf("%s", resp.GetMsg()), resp.GetMsg())
	}
	return nil
}

// Emit tirggers the initialize endpoint for the AppConfigSerivce asking it to
// initialize and acknowlege that the owners app permissions must be append by the created app
func (client ClientUserAuth) Emit(ctx context.Context, event *ports.Event, errC chan error) {
	resp, err := client.Conn.AddAppAccess(ctx, &grpcUserAuth.AddAppAccessRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		UserUuid:   event.App.OwnerUuid,
		AppUuid:    event.App.Uuid,
		AppRole:    *common.AppRole_OWNER.Enum(),
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("userpermission.addAppAccess err: %s", resp.GetMsg())
		return
	}
	errC <- nil
}
