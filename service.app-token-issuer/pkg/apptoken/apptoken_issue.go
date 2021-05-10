package apptoken

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/errors"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/storage"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/tokenissuer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (apt apptoken) Issue(ctx context.Context, db storage.Storage, claims TokenClaims) (*MetaToken, errors.ErrApi) {

	var storedClaims *TokenClaims = &TokenClaims{}
	err := db.FindOne(ctx, apptokenDB, apptokenColl, bson.M{"_id": claims.AppUuid}, storedClaims)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not Issue App-Token",
			Err:    err,
		}
	}
	// if storedClaims - check if can be overwritten - then issue token and update data
	var newExp = time.Now().Add(apptokenExp)
	if storedClaims.AppToken != "" {
		overrideOK := apt.canBeOverriden(storedClaims.Exp)
		if !overrideOK {
			return nil, errors.ErrAPI{
				Status: http.StatusBadRequest,
				Msg:    "Current App-Token is still valid",
				Err:    fmt.Errorf("app token must have expired before creating a new one"),
			}
		}
		storedClaims.Exp = newExp
		token, err := tokenissuer.IssueNew(claims.AppUuid, claims.AppHash, claims.AppOrigin, claims.Exp)
		if err != nil {
			return nil, errors.ErrAPI{
				Status: http.StatusInternalServerError,
				Msg:    "Could not issue App-Token",
				Err:    err,
			}
		}
		// upsert token data (token and exp)
		_, err = db.UpdateOne(ctx, apptokenDB, apptokenColl, bson.M{"_id": storedClaims.AppUuid}, bson.D{
			{Key: "$set", Value: bson.M{"app_token": token, "token_exp": newExp}}}, false)
		if err != nil {
			return nil, errors.ErrAPI{
				Status: http.StatusInternalServerError,
				Msg:    "Could not issue App-Token",
				Err:    err,
			}
		}
		return &MetaToken{Token: token, Exp: storedClaims.Exp.Unix()}, nil
	}
	// if no app token present issue token right away
	token, err := tokenissuer.IssueNew(claims.AppUuid, claims.AppHash, claims.AppOrigin, newExp)
	if err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not issue App-Token",
			Err:    err,
		}
	}
	claims.AppToken = token
	claims.Exp = newExp
	if err := db.InsertOne(ctx, apptokenDB, apptokenColl, claims); err != nil {
		return nil, errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not issue App-Token",
			Err:    err,
		}
	}
	return &MetaToken{Token: token, Exp: claims.Exp.Unix()}, nil
}
