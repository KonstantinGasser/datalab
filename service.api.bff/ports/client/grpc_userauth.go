package client

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
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

func (client ClientUserAuth) Register(ctx context.Context, r *users.RegisterRequest) (string, errors.Api) {
	resp, err := client.Conn.Register(ctx, &grpcUserAuth.RegisterRequest{
		Tracing_ID:   ctx_value.GetString(ctx, "tracingID"),
		Username:     r.Username,
		Organisation: r.Organization,
		Password:     r.Password,
	})
	if err != nil {
		return "", errors.New(http.StatusInternalServerError,
			err,
			"Could not create User Account")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return "", errors.New(resp.GetStatusCode(),
			err,
			resp.GetMsg())
	}
	return resp.GetUserUuid(), nil
}

func (client ClientUserAuth) Login(ctx context.Context, r *users.LoginRequest) (string, errors.Api) {
	resp, err := client.Conn.Login(ctx, &grpcUserAuth.LoginRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		Username:   r.Username,
		Password:   r.Password,
	})
	if err != nil {
		return "", errors.New(http.StatusInternalServerError,
			err,
			"Could not login User")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return "", errors.New(resp.GetStatusCode(),
			err,
			resp.GetMsg())
	}
	return resp.GetAccessToken(), nil
}

func (client ClientUserAuth) Authenticate(ctx context.Context, accessToken string) (*common.AuthedUser, errors.Api) {
	resp, err := client.Conn.IsAuthed(ctx, &grpcUserAuth.IsAuthedRequest{
		Tracing_ID:  ctx_value.GetString(ctx, "tracingID"),
		AccessToken: accessToken,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError,
			err,
			"Could authenticate User")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.New(resp.GetStatusCode(),
			err,
			resp.GetMsg())
	}
	return resp.GetAuthedUser(), nil
}
