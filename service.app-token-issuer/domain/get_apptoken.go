package domain

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

func (svc apptokenissuer) getAppToken(ctx context.Context, uuid string) (*common.AppTokenInfo, errors.ErrApi) {
	var token AppToken
	err := svc.repo.GetById(ctx, uuid, &token)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusBadRequest, err, "Could not find any data")
		}
		return nil, errors.New(http.StatusInternalServerError, err, "Could not find any data")
	}
	return &common.AppTokenInfo{Token: token.AppToken, Exp: token.Exp.Unix()}, nil
}
