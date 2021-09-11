package client

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	grpcAppToken "github.com/KonstantinGasser/datalab/service.app.token.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

type ClientAppToken struct {
	Conn grpcAppToken.AppTokenClient
}

func NewClientAppToken(clientAddr string) (*ClientAppToken, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := grpcAppToken.NewAppTokenClient(conn)
	return &ClientAppToken{
		Conn: client,
	}, nil
}

func (client ClientAppToken) Validate(ctx context.Context, token string) (string, string, errors.Api) {

	resp, err := client.Conn.Validate(ctx, &grpcAppToken.ValidateRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		AuthedUser: nil,
		AppToken:   token,
	})
	if err != nil {
		return "", "", errors.New(http.StatusInternalServerError,
			err,
			"Could not validate App Token")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return "", "", errors.New(resp.GetStatusCode(),
			err,
			resp.GetMsg())
	}
	return resp.GetAppUuid(), resp.GetAppOrigin(), nil
}
