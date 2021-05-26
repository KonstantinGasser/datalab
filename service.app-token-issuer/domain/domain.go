package domain

import (
	"context"
	"net/http"
	"time"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/permissions"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/KonstantinGasser/required"
)

const (
	tokenExp = time.Hour * 24 * 7
)

type AppTokenIssuer interface {
	InitToken(ctx context.Context, in *proto.InitRequest) errors.ErrApi
	IssueToken(ctx context.Context, in *proto.IssueRequest) (*common.AppTokenInfo, errors.ErrApi)
	ValidateToken(ctx context.Context, in *proto.ValidateRequest) (string, string, errors.ErrApi)
	GetToken(ctx context.Context, in *proto.GetRequest) (*common.AppTokenInfo, errors.ErrApi)
}

type apptokenissuer struct {
	dao        Dao
	permission permissions.Permission
}

func NewDomainLogic(dao Dao) AppTokenIssuer {
	return &apptokenissuer{
		dao:        dao,
		permission: permissions.New(dao),
	}
}

// InitToken initializes the doucment with the required meta data for an App Token
func (svc apptokenissuer) InitToken(ctx context.Context, in *proto.InitRequest) errors.ErrApi {
	var appToken = AppToken{
		AppUuid:  in.GetAppUuid(),
		AppHash:  in.GetAppHash(),
		AppOwner: in.GetAppOwner(),
	}
	if err := required.Atomic(&appToken); err != nil {
		return errors.New(http.StatusBadRequest, err, "Missing fields")
	}
	return svc.initAppToken(ctx, appToken)
}

// IssueToken creates a new app token and updates the token in the database if the token is not set yet
// or has expired else will return an error
func (svc apptokenissuer) IssueToken(ctx context.Context, in *proto.IssueRequest) (*common.AppTokenInfo, errors.ErrApi) {
	permissionErr := svc.permission.IsOwner(ctx, svc.dao, in.GetCallerUuid(), in.GetAppUuid())
	if permissionErr != nil {
		return nil, permissionErr
	}

	jwt, err := svc.issueAppToken(ctx, in.GetAppUuid(), in.GetAppOrigin(), in.GetAppHash())
	if err != nil {
		return nil, err
	}
	return jwt, nil
}

// ValidateToken checks if the token is a valid JWT and returns a subset of the data if valid
func (svc apptokenissuer) ValidateToken(ctx context.Context, in *proto.ValidateRequest) (string, string, errors.ErrApi) {
	appUuid, appOrigin, err := svc.validateAppToken(ctx, in.GetAppToken())
	if err != nil {
		return "", "", err
	}
	return appUuid, appOrigin, nil
}

// GetToken looks up the stored App Token data
func (svc apptokenissuer) GetToken(ctx context.Context, in *proto.GetRequest) (*common.AppTokenInfo, errors.ErrApi) {
	token, err := svc.getAppToken(ctx, in.GetAppUuid())
	if err != nil {
		return nil, err
	}
	return token, nil
}
