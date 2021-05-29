package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/library/errors"
	"github.com/KonstantinGasser/datalab/service.user-administer/domain/create"
	"github.com/KonstantinGasser/datalab/service.user-administer/domain/get"
	"github.com/KonstantinGasser/datalab/service.user-administer/domain/update"
	"github.com/KonstantinGasser/datalab/service.user-administer/proto"
	"github.com/KonstantinGasser/datalab/service.user-administer/repo"
)

type UserAdminLogic interface {
	CreateUser(ctx context.Context, in *proto.CreateRequest) errors.Api
	GetUsers(ctx context.Context, in *proto.GetListRequest) ([]*common.UserInfo, errors.Api)
	GetUser(ctx context.Context, in *proto.GetRequest) (*common.UserInfo, errors.Api)
	GetColleagues(ctx context.Context, in *proto.GetColleaguesRequest) ([]*common.UserInfo, errors.Api)
	UpdateUser(ctx context.Context, in *proto.UpdateRequest) errors.Api
}

type useradminlogic struct {
	repo repo.Repo
}

func NewUserAdminLogic(repo repo.Repo) UserAdminLogic {
	return &useradminlogic{
		repo: repo,
	}
}

func (svc useradminlogic) CreateUser(ctx context.Context, in *proto.CreateRequest) errors.Api {

	err := create.User(ctx, svc.repo, in)
	if err != nil {
		if err == create.ErrInvalidOrgnName {
			return errors.API{
				Status: http.StatusBadRequest,
				Msg:    "Provided Organization-Name must not include a forward-slash",
				Err:    err,
			}
		}
		if err == create.ErrUserNameTaken {
			return errors.API{
				Status: http.StatusBadRequest,
				Msg:    "Provided Username is already taken",
				Err:    err,
			}
		}
		return errors.API{
			Status: http.StatusInternalServerError,
			Msg:    "Could not create User-Account",
			Err:    err,
		}
	}
	return nil
}

func (svc useradminlogic) GetUser(ctx context.Context, in *proto.GetRequest) (*common.UserInfo, errors.Api) {
	user, err := get.User(ctx, svc.repo, in)
	if err != nil {
		if err == get.ErrNoUserFound {
			return nil, errors.API{
				Status: http.StatusNotFound,
				Msg:    "Could not find user",
				Err:    err,
			}
		}
		return nil, errors.API{
			Status: http.StatusInternalServerError,
			Msg:    "Could not find user",
			Err:    err,
		}
	}
	return user, nil
}

func (svc useradminlogic) GetUsers(ctx context.Context, in *proto.GetListRequest) ([]*common.UserInfo, errors.Api) {
	users, err := get.Users(ctx, svc.repo, in)
	if err != nil {
		if err == get.ErrNoUsersFound {
			return nil, errors.API{
				Status: http.StatusNotFound,
				Msg:    "Could not find any users",
				Err:    err,
			}
		}
		return nil, errors.API{
			Status: http.StatusInternalServerError,
			Msg:    "Could not find  users",
			Err:    err,
		}
	}
	return users, nil
}

func (svc useradminlogic) GetColleagues(ctx context.Context, in *proto.GetColleaguesRequest) ([]*common.UserInfo, errors.Api) {
	users, err := get.Colleaues(ctx, svc.repo, in.GetUserUuid())
	if err != nil {
		if err == get.ErrNoUserFound {
			return nil, errors.API{
				Status: http.StatusNotFound,
				Msg:    "Could not find requested users",
				Err:    err,
			}
		}
		if err == get.ErrNoUsersFound {
			return nil, errors.API{
				Status: http.StatusNotFound,
				Msg:    "Could not find any users",
				Err:    err,
			}
		}
		return nil, errors.API{
			Status: http.StatusInternalServerError,
			Msg:    "Could not find  users",
			Err:    err,
		}
	}
	return users, nil
}

func (svc useradminlogic) UpdateUser(ctx context.Context, in *proto.UpdateRequest) errors.Api {
	err := update.User(ctx, svc.repo, in)
	if err != nil {
		return errors.API{
			Status: http.StatusInternalServerError,
			Msg:    "Could not update user",
			Err:    err,
		}
	}
	return nil
}
