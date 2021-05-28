package client

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
	grpcUserMeta "github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"google.golang.org/grpc"
)

type ClientUserMeta struct {
	Conn grpcUserMeta.UserAdministerClient
}

func NewClientUserMeta(clientAddr string) (*ClientUserMeta, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := grpcUserMeta.NewUserAdministerClient(conn)
	return &ClientUserMeta{
		Conn: client,
	}, nil
}

func (client ClientUserMeta) CreateUserProfile(ctx context.Context, r *users.RegisterRequest) errors.Api {

	resp, err := client.Conn.Create(ctx, &grpcUserMeta.CreateRequest{
		Tracing_ID: ctx.Value("tracingID").(string),
		User: &common.UserInfo{
			Uuid:          r.UserUuid,
			Username:      r.Username,
			FirstName:     r.FirstName,
			LastName:      r.LastName,
			OrgnDomain:    r.Organization,
			OrgnPosition:  r.Position,
			ProfileImgUrl: "",
		},
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not create User Account")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			err,
			resp.GetMsg())
	}
	return nil
}

func (client ClientUserMeta) UpdateUserProfile(ctx context.Context, r *users.UpdateProfileRequest) errors.Api {

	resp, err := client.Conn.Update(ctx, &grpcUserMeta.UpdateRequest{
		Tracing_ID: ctx.Value("tracingID").(string),
		CallerUuid: r.UserUuid,
		User: &grpcUserMeta.UpdatableUser{
			FirstName:     r.FirstName,
			LastName:      r.LastName,
			OrgnPosition:  r.Organization,
			ProfileImgUrl: "",
		},
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not update User Profile")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			err,
			resp.GetMsg())
	}
	return nil
}
