package client

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
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

func (client ClientAppMeta) CreateApp(ctx context.Context, r *apps.CreateAppRequest) errors.Api {

	resp, err := client.Conn.Create(ctx, &grpcAppMeta.CreateRequest{
		Tracing_ID:   ctx.Value("tracingID").(string),
		OwnerUuid:    r.OwnerUuid,
		Name:         r.AppName,
		Organization: r.Organization,
		Description:  r.AppDesc,
		AppUrl:       r.AppUrl,
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not create App")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			err,
			resp.GetMsg())
	}
	return nil
}

func (client ClientAppMeta) GetApp(ctx context.Context, r *apps.GetAppRequest) (*common.AppInfo, errors.Api) {

	resp, err := client.Conn.Get(ctx, &grpcAppMeta.GetRequest{
		Tracing_ID: ctx.Value("tracingID").(string),
		AuthedUser: r.AuthedUser,
		AppUuid:    r.AppUuid,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError,
			err,
			"Could not create App")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.New(resp.GetStatusCode(),
			err,
			resp.GetMsg())
	}
	return resp.GetApp(), nil
}
