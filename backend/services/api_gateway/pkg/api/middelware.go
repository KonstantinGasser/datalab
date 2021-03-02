package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	tokenSrv "github.com/KonstantinGasser/clickstream/backend/grpc_definitions/token_service"
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
			logrus.Warnf("[api.WithAuth] %s is not authenticated", r.Host)
			api.onError(w, errors.New("no Authentication-Header found"), http.StatusForbidden)
			return
		}
		// invoke grpc call to token-service to validate a JWT
		ctx, cancel := context.WithTimeout(r.Context(), authTimeout)
		defer cancel()
		resp, err := api.TokenSrvClient.ValidateJWT(ctx, &tokenSrv.ValidateJWTRequest{
			JwtToken: token,
		})
		if err != nil {
			logrus.Errorf("[api.WithAuth] could not execute grpc.ValidateJWT: %v", err)
			api.onError(w, errors.New("no Authentication-Header found"), http.StatusInternalServerError)
			return
		}
		if resp.GetStatusCode() != http.StatusOK {
			logrus.Errorf("[api.WithAuth] grpc.ValidateJWT returned a %d code", resp.GetStatusCode())
			api.onError(w, fmt.Errorf("grpc received a code: %d", resp.GetStatusCode()), int(resp.GetStatusCode()))
			return
		}
		if !resp.GetIsValid() {
			logrus.Warnf("[api.WithAuth] %s is not authenticated: %v", r.Host, resp.GetMsg())
			api.onError(w, errors.New("not authenticated"), http.StatusForbidden)
			return
		}
		// serve request
		next(w, r)
	}
}
