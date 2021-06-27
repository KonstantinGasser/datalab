package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	grpcAppConfig "github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports"
	"github.com/KonstantinGasser/datalab/service.app.meta.agent/ports/client/intercepter"
	"google.golang.org/grpc"
)

type ClientAppConfig struct {
	conn grpcAppConfig.AppConfigurationClient
}

func NewClientAppConfig(clientAddr string) (*ClientAppConfig, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure(), intercepter.WithUnary(intercepter.WithAuth))
	if err != nil {
		return nil, err
	}
	client := grpcAppConfig.NewAppConfigurationClient(conn)
	return &ClientAppConfig{
		conn: client,
	}, nil
}

// EmitInit tirggers the initialize endpoint for the AppConfigSerivce asking it to
// initialize and acknowlege that a new app config with the given reference needs to be
// created
func (client ClientAppConfig) EmitInit(ctx context.Context, event *ports.InitEvent, errC chan error) {
	resp, err := client.conn.Initialize(ctx, &grpcAppConfig.InitRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppRefUuid: event.App.Uuid,
		AppOwner:   event.App.OwnerUuid,
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("appconf.init err: %s", resp.GetMsg())
		return
	}
	errC <- nil
}

// EmitAppendPermissions tirggers the add permissions endpoint for the AppConfigSerivce asking it to
// append its app config permissions with the new user joining the application
func (client ClientAppConfig) EmitAppendPermissions(ctx context.Context, event *ports.PermissionEvent, errC chan error) {
	resp, err := client.conn.AppendPermission(ctx, &grpcAppConfig.AppendPermissionRequest{
		Tracing_ID: "1",
		AppUuid:    event.AppUuid,
		UserUuid:   event.UserUuid,
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("apptoken.AppendPermisson err: %s", resp.GetMsg())
		return
	}
	errC <- nil
}

// EmitAppendPermissions tirggers the add permissions endpoint for the AppTokenSerivce asking it to
// add the joined user permissions to the app token
func (client ClientAppConfig) EmitRollbackAppendPermissions(ctx context.Context, event *ports.PermissionEvent, errC chan error) {
	resp, err := client.conn.RollbackAppendPermission(ctx, &grpcAppConfig.RollbackAppendPermissionRequest{
		Tracing_ID: "1",
		AppUuid:    event.AppUuid,
		UserUuid:   event.UserUuid,
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("apptoken.AppendPermisson err: %s", resp.GetMsg())
		return
	}
	errC <- nil
}
