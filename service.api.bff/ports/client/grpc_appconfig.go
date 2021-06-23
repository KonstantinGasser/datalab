package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/apps"
	"github.com/KonstantinGasser/datalab/service.api.bff/ports/client/intercepter"
	grpcAppConfig "github.com/KonstantinGasser/datalab/service.app.config.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

type ClientAppConfig struct {
	Conn grpcAppConfig.AppConfigurationClient
}

func NewClientAppConfig(clientAddr string) (*ClientAppConfig, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure(), intercepter.WithUnary(intercepter.WithAuth))
	if err != nil {
		return nil, err
	}
	client := grpcAppConfig.NewAppConfigurationClient(conn)
	return &ClientAppConfig{
		Conn: client,
	}, nil
}

func (client ClientAppConfig) CollectAppConfig(ctx context.Context, appUuid string, resC chan struct {
	Field string
	Value interface{}
}, errC chan error) {
	resp, err := client.Conn.Get(ctx, &grpcAppConfig.GetRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppUuid:    appUuid,
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("appconfig.get err: %s", resp.GetMsg())
		return
	}
	resC <- struct {
		Field string
		Value interface{}
	}{
		Field: "appconfig",
		Value: resp.GetConfigs(),
	}
}

func (client ClientAppConfig) UpdateConfig(ctx context.Context, r *apps.UpdateConfigRequest) errors.Api {
	resp, err := client.Conn.Update(ctx, &grpcAppConfig.UpdateRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppRefUuid: r.AppRefUuid,
		UpdateFlag: r.UpdateFlag,
		Stages:     r.Stages,
		Records:    r.Records,
		BtnDefs:    r.BtnDefs,
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not update App Config")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return nil
}

func (client ClientAppConfig) LockAppConfig(ctx context.Context, appUuid string, authedUser *common.AuthedUser) errors.Api {
	resp, err := client.Conn.LockConfig(ctx, &grpcAppConfig.LockConfigRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppRefUuid: appUuid,
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not lock app config")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return nil
}

func (client ClientAppConfig) UnlockAppConfig(ctx context.Context, r *apps.UnlockRequest) errors.Api {
	resp, err := client.Conn.UnlockConfig(ctx, &grpcAppConfig.UnlockConfigRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppRefUuid: r.AppUuid,
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not unlock app config")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return nil
}
