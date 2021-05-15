package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.user-permissions/domain/initialize"
	"github.com/KonstantinGasser/datalab/service.user-permissions/errors"
	"github.com/KonstantinGasser/datalab/service.user-permissions/proto"
	"github.com/KonstantinGasser/datalab/service.user-permissions/repo"
)

type UserPermissions interface {
	InitPermissions(ctx context.Context, in *proto.InitRequest) errors.ErrApi
}

type userpermissions struct {
	repo repo.Repo
}

func NewPermissionsLogic(repo repo.Repo) UserPermissions {
	return &userpermissions{
		repo: repo,
	}
}

func (svc userpermissions) InitPermissions(ctx context.Context, in *proto.InitRequest) errors.ErrApi {

	err := initialize.Permissions(ctx, svc.repo, in)
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not initialize User Permissions",
			Err:    err,
		}
	}
	return nil
}
