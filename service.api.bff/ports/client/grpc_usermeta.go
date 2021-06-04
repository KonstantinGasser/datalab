package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.api.bff/internal/users"
	grpcUserMeta "github.com/KonstantinGasser/datalab/service.user.meta.agent/cmd/grpcserver/proto"
	"google.golang.org/grpc"
)

type ClientUserMeta struct {
	Conn grpcUserMeta.UserMetaClient
}

func NewClientUserMeta(clientAddr string) (*ClientUserMeta, error) {
	conn, err := grpc.Dial(clientAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := grpcUserMeta.NewUserMetaClient(conn)
	return &ClientUserMeta{
		Conn: client,
	}, nil
}

func (client ClientUserMeta) CreateUserProfile(ctx context.Context, r *users.RegisterRequest) errors.Api {

	resp, err := client.Conn.Create(ctx, &grpcUserMeta.CreateRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		User: &common.UserInfo{
			Uuid:         r.UserUuid,
			Username:     r.Username,
			FirstName:    r.FirstName,
			LastName:     r.LastName,
			OrgnDomain:   r.Organization,
			OrgnPosition: r.Position,
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
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		CallerUuid: r.UserUuid,
		User: &grpcUserMeta.UpdatableUser{
			FirstName:    r.FirstName,
			LastName:     r.LastName,
			OrgnPosition: r.Position,
		},
	})
	if err != nil {
		return errors.New(http.StatusInternalServerError,
			err,
			"Could not update User Profile")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return nil
}

func (client ClientUserMeta) GetProfile(ctx context.Context, r *users.GetProfileRequest) (*common.UserInfo, errors.Api) {

	resp, err := client.Conn.Get(ctx, &grpcUserMeta.GetRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		CallerUuid: r.UserUuid,
		ForUuid:    r.UserUuid,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError,
			err,
			"Could not get User Profile")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return resp.GetUser(), nil
}

func (client ClientUserMeta) GetColleagues(ctx context.Context, r *users.GetColleagueRequest) ([]*common.UserInfo, errors.Api) {

	resp, err := client.Conn.GetColleagues(ctx, &grpcUserMeta.GetColleaguesRequest{
		Tracing_ID:   ctx_value.GetString(ctx, "tracingID"),
		Organization: r.Organization,
	})
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError,
			err,
			"Could not get User Profile")
	}
	if resp.GetStatusCode() != http.StatusOK {
		return nil, errors.New(resp.GetStatusCode(),
			fmt.Errorf(resp.GetMsg()),
			resp.GetMsg())
	}
	return resp.GetColleagues(), nil
}

func (client ClientUserMeta) CollectOwnerInfo(ctx context.Context, appOwner string, resC chan struct {
	Field string
	Value interface{}
}, errC chan error) {
	resp, err := client.Conn.Get(ctx, &grpcUserMeta.GetRequest{
		Tracing_ID: ctx_value.GetString(ctx, "tracingID"),
		CallerUuid: appOwner,
		ForUuid:    appOwner,
	})
	if err != nil {
		errC <- err
		return
	}
	if resp.GetStatusCode() != http.StatusOK {
		errC <- fmt.Errorf("usermeta.get err: %s", resp.GetMsg())
		return
	}
	resC <- struct {
		Field string
		Value interface{}
	}{
		Field: "appowner",
		Value: resp.GetUser(),
	}
}
