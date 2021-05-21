package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-administer/errors"
	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/get"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/initialize"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/issue"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/permissions"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/validate"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/repo"
)

type AppTokenIssuer interface {
	InitToken(ctx context.Context, in *proto.InitRequest) errors.ErrApi
	IssueToken(ctx context.Context, in *proto.IssueRequest) (*common.AppTokenInfo, errors.ErrApi)
	ValidateToken(ctx context.Context, in *proto.ValidateRequest) (string, string, errors.ErrApi)
	GetToken(ctx context.Context, in *proto.GetRequest) (*common.AppTokenInfo, errors.ErrApi)
}

type apptokenissuer struct {
	repo     repo.Repo
	appadmin appsvc.AppAdministerClient
}

func NewAppTokenLogic(repo repo.Repo, appadmin appsvc.AppAdministerClient) AppTokenIssuer {
	return &apptokenissuer{
		repo:     repo,
		appadmin: appadmin,
	}
}

func (svc apptokenissuer) InitToken(ctx context.Context, in *proto.InitRequest) errors.ErrApi {
	err := initialize.AppToken(ctx, svc.repo, in.GetAppUuid(), in.GetAppHash(), in.GetAppOwner())
	if err != nil {
		return errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not initialize App Token",
			Err:    err,
		}
	}
	return nil
}

func (svc apptokenissuer) IssueToken(ctx context.Context, in *proto.IssueRequest) (*common.AppTokenInfo, errors.ErrApi) {
	permissionErr := permissions.IsOwner(ctx, svc.repo, in.GetCallerUuid(), in.GetAppUuid())
	if permissionErr != nil {
		if permissionErr == permissions.ErrNotAuthorized {
			return nil, errors.ErrAPI{
				Status: http.StatusUnauthorized,
				Msg:    "User must be App Owner to generate App Token",
				Err:    permissionErr,
			}
		}
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not verify permissions",
			Err:    permissionErr,
		}
	}

	token, err := issue.Token(ctx, svc.repo, in)
	if err != nil {
		if err == issue.ErrTokenStillValid {
			return nil, errors.ErrAPI{
				Status: http.StatusBadRequest,
				Msg:    "Current App-Token has not expired yet",
				Err:    err,
			}
		}
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not issue App-Token",
			Err:    err,
		}
	}
	return token, nil
}

func (svc apptokenissuer) ValidateToken(ctx context.Context, in *proto.ValidateRequest) (string, string, errors.ErrApi) {
	appUuid, appOrigin, err := validate.Token(ctx, in.GetAppToken())
	if err != nil {
		return "", "", errors.ErrAPI{
			Status: http.StatusUnauthorized,
			Msg:    "App Token invalid",
			Err:    err,
		}
	}
	return appUuid, appOrigin, nil
}

func (svc apptokenissuer) GetToken(ctx context.Context, in *proto.GetRequest) (*common.AppTokenInfo, errors.ErrApi) {
	permissionErr := permissions.CanAccess(ctx, svc.repo, in.GetUserClaims(), in.GetAppUuid())
	if permissionErr != nil {
		if permissionErr == permissions.ErrNotAuthorized {
			return nil, errors.ErrAPI{
				Status: http.StatusUnauthorized,
				Msg:    "User must be App Owner to generate App Token",
				Err:    permissionErr,
			}
		}
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not verify permissions",
			Err:    permissionErr,
		}
	}
	token, err := get.Token(ctx, svc.repo, in.GetAppUuid())
	if err != nil {
		if err == get.ErrNotFound {
			return nil, errors.ErrAPI{
				Status: http.StatusNotFound,
				Msg:    "Could not find App-Token",
				Err:    err,
			}
		}
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not find App-Token",
			Err:    err,
		}
	}
	return token, nil
}
