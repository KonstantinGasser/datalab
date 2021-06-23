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
	grpcAppToken "github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

type ClientAppToken struct {
	Conn grpcAppToken.AppTokenClient
}

func NewClientAppToken(clientAddr string) (*ClientAppToken, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure(), intercepter.WithUnary(intercepter.WithAuth))
	if err != nil {
		return nil, err
	}
	client := grpcAppToken.NewAppTokenClient(conn)
	return &ClientAppToken{
		Conn: client,
	}, nil
}

func (client ClientAppToken) CollectAppToken(ctx context.Context, appUuid string, resC chan struct {
	Field string
	Value interface{}
}, errC chan error) {
	resp, err := client.Conn.Get(ctx, &grpcAppToken.GetRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppUuid:    appUuid,
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("apptoken.get err: %s", resp.GetMsg())
		return
	}
	resC <- struct {
		Field string
		Value interface{}
	}{
		Field: "apptoken",
		Value: resp.GetToken(),
	}
}

func (client ClientAppToken) IssueAppToken(ctx context.Context, r *apps.CreateAppTokenRequest) (*common.AppAccessToken, errors.Api) {

	resp, err := client.Conn.Issue(ctx, &grpcAppToken.IssueRequest{
		Tracing_ID:   ctx_value.GetString(ctx, "tracingID"),
		AppUuid:      r.AppUuid,
		AppName:      r.AppName,
		Organization: r.Organization,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError,
			err,
			"Could not issue App Token")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.New(resp.GetStatusCode(),
			err,
			resp.GetMsg())
	}
	return resp.GetToken(), nil
}

func (client ClientAppToken) UnlockAppToken(ctx context.Context, r *apps.UnlockRequest) errors.Api {
	resp, err := client.Conn.UnlockAppToken(ctx, &grpcAppToken.UnlockAppTokenRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AppUuid:    r.AppUuid,
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not unlock App Token")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			err,
			resp.GetMsg())
	}
	return nil
}
