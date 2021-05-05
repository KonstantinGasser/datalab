package apptoken

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/storage"
	"github.com/KonstantinGasser/datalab/service_apptoken/pkg/tokenissuer"
	"github.com/KonstantinGasser/datalab/service_config/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (apt apptoken) Issue(ctx context.Context, db storage.Storage, claims TokenClaims) (string, errors.ErrApi) {

	ok, err := db.Exists(ctx, apptokenDB, apptokenDB, bson.M{"_id": claims.AppUuid})
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not Issue App-Token",
			Err:    err,
		}
	}
	if ok {
		overrideOK, err := apt.canBeOverriden(ctx, db, claims.AppUuid)
		if err != nil {
			return "", errors.ErrAPI{
				Status: http.StatusInternalServerError,
				Msg:    "Could not issue App-Token",
				Err:    err,
			}
		}
		if !overrideOK {
			return "", errors.ErrAPI{
				Status: http.StatusBadRequest,
				Msg:    "Current Token is still valid",
				Err:    fmt.Errorf("current app token is still valid, token must have expired before create a new one"),
			}
		}
	}

	expTime := time.Now().Add(apptokenExp)
	claims.Exp = expTime
	token, err := tokenissuer.IssueNew(claims.AppUuid, claims.AppHash, claims.AppOrigin, claims.Exp)
	if err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not issue App-Token",
			Err:    err,
		}
	}
	claims.AppToken = token
	if err := db.InsertOne(ctx, apptokenDB, apptokenColl, claims); err != nil {
		return "", errors.ErrAPI{
			Status: http.StatusInternalServerError,
			Msg:    "Could not issue App-Token",
			Err:    err,
		}
	}
	return token, nil
}
