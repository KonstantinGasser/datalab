package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-administer/errors"
	appsvc "github.com/KonstantinGasser/datalab/service.app-administer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/get"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/issue"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/repo"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppTokenIssuer interface {
	IssueToken(ctx context.Context, in *proto.IssueRequest) (*common.AppTokenInfo, errors.ErrApi)
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

func (svc apptokenissuer) IssueToken(ctx context.Context, in *proto.IssueRequest) (*common.AppTokenInfo, errors.ErrApi) {
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

func (svc apptokenissuer) GetToken(ctx context.Context, in *proto.GetRequest) (*common.AppTokenInfo, errors.ErrApi) {
	token, err := get.Token(ctx, svc.repo, in.GetAppUuid())
	if err != nil {
		if err == mongo.ErrNoDocuments {
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
