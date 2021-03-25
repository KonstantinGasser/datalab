package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
	"github.com/KonstantinGasser/clickstream/utils/ctx_value"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

const (
	authTimeout = time.Second * 5
)

// WithCors enables CORS by setting the 'Access-Control-Allow-Origin' and
// 'Access-Control-Allow-Methods' header as specified by the API struct
func (api API) WithCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// loop over CORS-Config and set headers
		for _, item := range api.cors.Cfgs {
			w.Header().Set(item.Header, item.Value)
		}
		// return status OK if requested method is OPTIONS
		// (used to do pre-flights from browser with POST request)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// serve request
		next(w, r)
	}
}

// WithAuth acts as middleware before route endpoints. It checks if a user is authenticated or not
// returns a new http.HandlerFunc to allow multiple other middleware to be wrapped around/after
func (api API) WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := api.headerAuth(r)
		if err != nil {
			logrus.Warnf("<%v>[api.WithAuth] %s is not authenticated", ctx_value.GetString(r.Context(), "tracingID"), r.Host)
			api.onError(w, errors.New("no Authentication-Header found"), http.StatusForbidden)
			return
		}
		// invoke grpc call to token-service to validate a JWT
		// ctx := context.WithTimeout(r.Context(), authTimeout)
		// defer cancel()
		resp, err := api.TokenSrvClient.ValidateJWT(r.Context(), &tokenSrv.ValidateJWTRequest{
			JwtToken:   token,
			Tracing_ID: ctx_value.GetString(r.Context(), "tracingID"),
		})
		if err != nil {
			logrus.Errorf("<%v>[api.WithAuth] could not execute grpc.ValidateJWT: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
			api.onError(w, errors.New("could not execute authentication"), http.StatusInternalServerError)
			return
		}
		if resp.GetStatusCode() != http.StatusOK || !resp.GetIsValid() || resp.GetUser() == nil {
			logrus.Warnf("<%v>[api.WithAuth] %s is not authenticated: %v", ctx_value.GetString(r.Context(), "tracingID"), r.Host, resp.GetMsg())
			api.onError(w, errors.New("not authenticated"), http.StatusUnauthorized)
			return
		}
		// add JWT claims of user in r.Context()
		ctxWithVal := ctx_value.AddValue(r.Context(), "user", resp.GetUser())
		// serve request with user claims in context
		next(w, r.WithContext(ctxWithVal))
	}
}

// WithTracing allows to generate a tracing ID at the entry-point of an request which gets passed
// in the request.Context in order for it to be available in following code.
// The tracing ID is an straight forward approach to trace logs from multiple services
// Tracing ID are based on the current time and the MAC-Address
func (api API) WithTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tracingID, err := uuid.NewV1()
		if err != nil {
			// in case creating of tracing ID fails - don't border but server request!
			logrus.Infof("[api.WithTracing] could not create tracing ID: %v\n", err)
			next(w, r)
			return
		}
		// add tracing ID to request context for other function involved in the request
		// to have access to it
		ctx := ctx_value.AddValue(r.Context(), "tracingID", fmt.Sprintf("%x", tracingID.Bytes()[:4]))
		next(w, r.WithContext(ctx))
	}
}
