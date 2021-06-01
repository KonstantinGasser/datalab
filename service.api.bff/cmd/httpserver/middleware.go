package httpserver

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
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
func (s *Server) WithCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set CORS-Header to response writer as defined by the api
		w.Header().Set(accessControlAllowOrigin, strings.Join(s.allowedOrigins, ","))
		w.Header().Set(accessControlAllowMethods, strings.Join(s.allowedMethods, ","))
		w.Header().Set(accessControlAllowHeader, strings.Join(s.allowedHeaders, ","))

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// serve next handler
		next(w, r)
	}
}

// WithAuth acts as middleware before route endpoints. It checks if a user is authenticated or not
// returns a new http.HandlerFunc to allow multiple other middleware to be wrapped around/after
func (s *Server) WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")
		if accessToken == "" {
			logrus.Warnf("[%v][handler.WithAuth] %s is not authenticated", ctx_value.GetString(r.Context(), "tracingID"), r.Host)
			s.onErr(w, http.StatusForbidden, "missing accesss token")
			return
		}
		authedUser, err := s.userauthService.Authenticate(r.Context(), accessToken)
		if err != nil {
			logrus.Errorf("[%v][Server.WithAuth] could not authenticate user: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
			s.onErr(w, err.Code(), err.Info())
			return
		}
		// add JWT claims of user in r.Context()
		ctxWithVal := ctx_value.AddValue(r.Context(), "user", authedUser)
		// serve request with user claims in context
		next(w, r.WithContext(ctxWithVal))
	}
}

// WithTracing allows to generate a tracing ID at the entry-point of an request which gets added
// in the request.Context in order for it to be available through out the code.
// The tracing ID is an straight forward approach to trace logs from multiple services
// The Tracing ID is based on the current time and the MAC-Address
func (s Server) WithTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tracingID, err := uuid.NewV1()
		if err != nil {
			// in case creating of tracing ID fails - don't border but server request!
			logrus.Infof("[Server.WithTracing] could not create tracing ID: %v\n", err)
			next(w, r)
			return
		}
		// add tracing ID to request context for other function involved in the request
		// to have access to it
		ctx := ctx_value.AddValue(r.Context(), "tracingID", fmt.Sprintf("%x", tracingID.Bytes()[:4]))
		next(w, r.WithContext(ctx))
	}
}
