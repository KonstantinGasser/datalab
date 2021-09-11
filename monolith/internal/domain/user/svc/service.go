package svc

import (
	"context"

	"github.com/KonstantinGasser/datalab/monolith/internal/domain/user"
	"github.com/KonstantinGasser/datalab/monolith/pkg/ctxkey"
	"github.com/KonstantinGasser/datalab/monolith/pkg/httperr"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, username, firstname, lastname, organization, position string) httperr.ErrorHTTP
	Update(ctx context.Context) httperr.ErrorHTTP
	Details(ctx context.Context) (*user.User, httperr.ErrorHTTP)
	ByID(ctx context.Context, UUID uuid.UUID) (*user.User, httperr.ErrorHTTP)
	Colleagues(ctx context.Context, org string) ([]*user.User, httperr.ErrorHTTP)
}

type usersvc struct {
	repo user.Repository
}

func New(repo user.Repository) Service {
	return &usersvc{repo: repo}
}

func (s usersvc) Create(ctx context.Context, username, firstname, lastname, organization, position string) httperr.ErrorHTTP {
	newUser, err := user.NewDefault(username, firstname, lastname, organization, position)
	if err != nil {
		return httperr.InternalServerError(err.Error(), err)
	}

	if err := s.repo.Store(ctx, newUser); err != nil {
		return httperr.InternalServerError(err.Error(), err)
	}
	return nil
}

func (s usersvc) Update(ctx context.Context) httperr.ErrorHTTP {
	return nil
}

func (s usersvc) Details(ctx context.Context) (*user.User, httperr.ErrorHTTP) {
	session, err := ctxkey.Session(ctx)
	if err != nil {
		return nil, httperr.Unauthorized(err.Error(), err)
	}
	var u user.User
	if err := s.repo.ById(ctx, session.UUID, &u); err != nil {
		if err == user.ErrNotFound {
			return nil, httperr.BadRequest("user not found", err)
		}
		return nil, httperr.InternalServerError(err.Error(), err)
	}
	return &u, nil
}

func (s usersvc) ByID(ctx context.Context, UUID uuid.UUID) (*user.User, httperr.ErrorHTTP) {
	var u user.User
	if err := s.repo.ById(ctx, UUID, &u); err != nil {
		if err == user.ErrNotFound {
			return nil, httperr.BadRequest("user not found", err)
		}
		return nil, httperr.InternalServerError(err.Error(), err)
	}
	return &u, nil
}

func (s usersvc) Colleagues(ctx context.Context, org string) ([]*user.User, httperr.ErrorHTTP) {
	var u []*user.User
	if err := s.repo.ByOrganization(ctx, org, u); err != nil {
		if err == user.ErrNotFound {
			return nil, httperr.BadRequest("user not found", err)
		}
		return nil, httperr.InternalServerError(err.Error(), err)
	}
	return u, nil
}
