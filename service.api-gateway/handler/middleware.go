package handler

import (
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// WithAuth acts as middleware before route endpoints. It checks if a user is authenticated or not
// returns a new http.HandlerFunc to allow multiple other middleware to be wrapped around/after
func (handler *Handler) WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			logrus.Warnf("<%v>[handler.WithAuth] %s is not authenticated", ctx_value.GetString(r.Context(), "tracingID"), r.Host)
			handler.onError(w, "no Authentication-Header found", http.StatusForbidden)
			return
		}
		// invoke grpc call to token-service to validate a JWT
		claims, err := handler.domain.IsLoggedIn(r.Context(), token)
		if err != nil {
			logrus.Errorf("<%v>[handler.WithAuth] could not authenticate user: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
			handler.onError(w, err.Info(), int(err.Code()))
			return
		}
		// add JWT claims of user in r.Context()
		ctxWithVal := ctx_value.AddValue(r.Context(), "user", claims)
		// serve request with user claims in context
		next(w, r.WithContext(ctxWithVal))
	}
}

// WithTracing allows to generate a tracing ID at the entry-point of an request which gets added
// in the request.Context in order for it to be available through out the code.
// The tracing ID is an straight forward approach to trace logs from multiple services
// The Tracing ID is based on the current time and the MAC-Address
func (handler *Handler) WithTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tracingID, err := uuid.NewV1()
		if err != nil {
			// in case creating of tracing ID fails - don't border but server request!
			logrus.Infof("[handler.WithTracing] could not create tracing ID: %v\n", err)
			next(w, r)
			return
		}
		// add tracing ID to request context for other function involved in the request
		// to have access to it
		ctx := ctx_value.AddValue(r.Context(), "tracingID", fmt.Sprintf("%x", tracingID.Bytes()[:4]))
		next(w, r.WithContext(ctx))
	}
}
