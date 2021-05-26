package domain

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/KonstantinGasser/datalab/common"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/domain/jwts"
	"github.com/KonstantinGasser/datalab/service.app-token-issuer/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrTokenStillValid = fmt.Errorf("app token is still valid. Thus can not be re-generated")
)

// issue AppToken updates the jwt and expiration time of an app token if their either is no jwt present
// or if the jwt has expired
func (svc apptokenissuer) issueAppToken(ctx context.Context, uuid, origin, hash string) (*common.AppTokenInfo, errors.ErrApi) {
	var stored *AppToken = &AppToken{}
	err := svc.dao.GetById(ctx, uuid, &stored)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New(http.StatusBadRequest, err, "No App Token found for updating")
		}
		return nil, errors.New(http.StatusInternalServerError, err, "Could not issue new App Token")
	}

	// logic for when app token already has been created
	// new app token can only be issued if current one has expired
	if stored.AppToken != "" {
		if ok := override(stored.Exp); !ok {
			return nil, errors.New(http.StatusBadRequest, ErrTokenStillValid, "Current App Token has not yet expired")
		}
	}

	var newExp = time.Now().Add(tokenExp)
	token, err := jwts.Generate(uuid, origin, hash, newExp)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, err, "Could not issue new Token")
	}
	if err := svc.dao.UpdateAppToken(ctx, uuid, token, origin, newExp); err != nil {
		return nil, errors.New(http.StatusInternalServerError, err, "Could not Update App Token")
	}
	return &common.AppTokenInfo{
		Token: token,
		Exp:   newExp.Unix(),
	}, nil
}

// override checks if a token has already expired or not
func override(exp time.Time) bool {
	return exp.Unix() <= time.Now().Unix()
}
