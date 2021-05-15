package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-configuration/errors"
	"github.com/KonstantinGasser/datalab/service.user-administer/domain/create"
	"github.com/KonstantinGasser/datalab/service.user-administer/domain/get"
	"github.com/KonstantinGasser/datalab/service.user-administer/domain/update"
	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/service.user-administer/repo"
	permissionssvc "github.com/KonstantinGasser/datalab/service.user-permissions/proto"
)

type UserAdminLogic interface {
	CreateUser(ctx context.Context, in *proto.CreateRequest) errors.ErrApi
	GetUsers(ctx context.Context, in *proto.GetListRequest) ([]*common.UserInfo, errors.ErrApi)
	GetUser(ctx context.Context, in *proto.GetRequest) (*common.UserInfo, errors.ErrApi)
	UpdateUser(ctx context.Context, in *proto.UpdateRequest) errors.ErrApi
}

type useradminlogic struct {
	permissions permissionssvc.UserPermissionsClient
	repo        repo.Repo
}

func NewUserAdminLogic(repo repo.Repo, permissions permissionssvc.UserPermissionsClient) UserAdminLogic {
	return &useradminlogic{
		permissions: permissions,
		repo:        repo,
	}
}

func (svc useradminlogic) CreateUser(ctx context.Context, in *proto.CreateRequest) errors.ErrApi {

	err := create.User(ctx, svc.repo, in)
	if err != nil {
		if err == create.ErrInvalidOrgnName {
			return errors.ErrAPI{
				Status: http.StatusBadRequest,
				Msg:    "Provided Organization-Name must not include a forward-slash",
				Err:    err,
			}
		}
		if err == create.ErrUserNameTaken {
			return errors.ErrAPI{
				Status: http.StatusBadRequest,
				Msg:    "Provided Username is already taken",
				Err:    err,
			}
		}
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create User-Account",
			Err:    err,
		}
	}
	resp, err := svc.permissions.Init(ctx, &permissionssvc.InitRequest{
		Tracing_ID: in.GetTracing_ID(),
		UserUuid:   in.GetUser().GetUuid(),
		UserOrgn:   in.GetUser().GetOrgnDomain()},
	)
	if err != nil {
		_ = create.Compansate(ctx, svc.repo, in.GetUser().GetUuid())
		return errors.ErrAPI{Status: http.StatusInternalServerError, Msg: "Could not init User Permissions", Err: err}
	}
	if resp.GetStatusCode() != http.StatusOK {
		_ = create.Compansate(ctx, svc.repo, in.GetUser().GetUuid())
		return errors.ErrAPI{Status: resp.GetStatusCode(), Msg: resp.GetMsg(), Err: fmt.Errorf("%s", resp.GetMsg())}
	}
	return nil
}

func (svc useradminlogic) GetUser(ctx context.Context, in *proto.GetRequest) (*common.UserInfo, errors.ErrApi) {
	user, err := get.User(ctx, svc.repo, in)
	if err != nil {
		if err == get.ErrNoUserFound {
			return nil, errors.ErrAPI{
				Status: http.StatusNotFound,
				Msg:    "Could not find user",
				Err:    err,
			}
		}
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not find user",
			Err:    err,
		}
	}
	return user, nil
}

func (svc useradminlogic) GetUsers(ctx context.Context, in *proto.GetListRequest) ([]*common.UserInfo, errors.ErrApi) {
	users, err := get.Users(ctx, svc.repo, in)
	if err != nil {
		if err == get.ErrNoUsersFound {
			return nil, errors.ErrAPI{
				Status: http.StatusNotFound,
				Msg:    "Could not find any users",
				Err:    err,
			}
		}
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not find  users",
			Err:    err,
		}
	}
	return users, nil
}

func (svc useradminlogic) UpdateUser(ctx context.Context, in *proto.UpdateRequest) errors.ErrApi {
	err := update.User(ctx, svc.repo, in)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not update user",
			Err:    err,
		}
	}
	return nil
}
