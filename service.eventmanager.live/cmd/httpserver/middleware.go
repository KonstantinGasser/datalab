package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/service.eventmanager.live/jwts"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// WithCors apply the CORS-Configs of the API to a given API-Endpoint
func (s *Server) WithCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set CORS-Header to response writer as defined by the api
		w.Header().Set(accessControlAllowOrigin, strings.Join(s.allowedOrigins, ","))
		w.Header().Set(accessControlAllowMethods, strings.Join(s.allowedMethods, ","))
		w.Header().Set(accessControlAllowHeaders, strings.Join(s.allowedHeaders, ","))

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
		accessToken := r.Header.Get(keyDatalabToken)
		if accessToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		appUuid, appOrigin, err := s.appTokenClient.Validate(r.Context(), accessToken)
		if err != nil {
			logrus.Errorf("[%v][Server.WithAuth] could not authenticate user: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
			s.onErr(w, err.Code(), err.Info())
			return
		}
		// add JWT claims of user in r.Context()
		ctxWithVal := context.WithValue(r.Context(), typeKeyClaims(keyAppUuid), appUuid)
		ctxWithVal = context.WithValue(ctxWithVal, typeKeyClaims(keyOrigin), appOrigin)
		// serve request with user claims in context
		next(w, r.WithContext(ctxWithVal))
	}
}

// WithTraceIP adds the requester's client IP to the request context
// serving as device identifier - if none present request is not accpeted
func (s *Server) WithTraceIP(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deviceIP := strings.Split(r.RemoteAddr, ":")
		if len(deviceIP) == 0 {
			s.onErr(w, http.StatusBadRequest, "not sufficiant information provided")
			return
		}
		ipCtx := context.WithValue(r.Context(), typeKeyIP(keyIP), deviceIP[1])

		next(w, r.WithContext(ipCtx))
	}
}

// WithTicketAuth looks for the from this service issued web-socket ticket
// and validates the ticket - if of it passed the request forward to the next caller
// else returns with a http error
func (s *Server) WithTicketAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ticket := r.URL.Query().Get("ticket")
		if ticket == "" {
			logrus.Errorf("could not find any ws-ticket want: jwt - have: %s\n", ticket)
			s.onErr(w, http.StatusUnauthorized, "no ws ticket found")
			return
		}
		claims, err := jwts.Validate(ticket)
		if err != nil {
			logrus.Error("provided ws-ticket is not valid")
			s.onErr(w, http.StatusUnauthorized, "not authorized")
			return
		}
		ctx := ctx_value.AddValue(r.Context(), keyOrigin, claims["origin"])
		next(w, r.WithContext(ctx))
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
