package middelware

import (
	"context"
	"net/http"

	"github.com/KonstantinGasser/datalab/monolith/pkg/ctxkey"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

const (
	// accessControlAllowOrigin describes the allowed request origins
	accessControlAllowOrigin = "Access-Control-Allow-Origin"
	// accessControlAllowMethods describes the methods allowed by this API
	accessControlAllowMethods = "Access-Control-Allow-Methods"
	// accessControlAllowHeader describes a header -> ???
	accessControlAllowHeader = "Access-Control-Allow-Headers"
)

// WithCors apply the CORS-Configs of the API to a given API-Endpoint
func WithCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set CORS-Header to response writer as defined by the api
		w.Header().Set(accessControlAllowOrigin, "127.0.0.1:3000")
		w.Header().Set(accessControlAllowMethods, "GET,POST")
		w.Header().Set(accessControlAllowHeader, "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// serve next handler
		next.ServeHTTP(w, r)
	})
}

// WithAuth acts as middleware before route endpoints. It checks if a user is authenticated or not
// returns a new http.HandlerFunc to allow multiple other middleware to be wrapped around/after
func WithAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// serve request with user claims in context
		next.ServeHTTP(w, r)
	})
}

// WithTracing allows to generate a tracing ID at the entry-point of an request which gets added
// in the request.Context in order for it to be available through out the code.
// The tracing ID is an straight forward approach to trace logs from multiple services
// The Tracing ID is based on the current time and the MAC-Address
func WithTracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tracingID, err := uuid.NewV1()
		if err != nil {
			// in case creating of tracing ID fails - don't border but server request!
			logrus.Infof("[Server.WithTracing] could not create tracing ID: %v\n", err)
			next.ServeHTTP(w, r)
			return
		}
		// add tracing ID to request context for other function involved in the request
		// to have access to it
		ctx := context.WithValue(r.Context(), ctxkey.Str("tracingID"), tracingID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Printf("[%s] -> incoming request from: %v\n",
			r.Context().Value(ctxkey.Str("tracingID")),
			r.RemoteAddr,
		)
		next.ServeHTTP(w, r)
	})
}
