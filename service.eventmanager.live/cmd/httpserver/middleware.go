package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/KonstantinGasser/datalab/library/utils/unique"
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
		ctxWithVal := ctx_value.AddValue(r.Context(), "app.uuid", appUuid)
		ctxWithVal = ctx_value.AddValue(ctxWithVal, "app.origin", appOrigin)
		// serve request with user claims in context
		next(w, r.WithContext(ctxWithVal))
	}
}

// WithCookie looks-up if a request already has an x-datalab cookie set
// else sets a new x-datalab cookie. In both cases the cookie information
// gets passed into the r.Context
func (s *Server) WithCookie(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cookie *http.Cookie
		var err error

		cookie, err = r.Cookie(keyCookie)
		if err != nil || cookie.Value == "" {
			logrus.Warnf("[middleware.WithCookie] no cookie present - setting new cookie\n")
			// set new cookie for request
			uuid, err := unique.UUID()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			cookie = &http.Cookie{
				Name:    keyCookie,
				Value:   uuid,
				Expires: time.Now().Add(1 * time.Hour),
				Path:    "/",
				Domain:  "sample.router.dev",
			}
			http.SetCookie(w, cookie)
		}
		// pass cookie via context
		ctx := context.WithValue(r.Context(), typeKeyCookie(keyCookie), cookie)
		// move to next handler
		next(w, r.WithContext(ctx))
	}
}

// MustCookie ensures that a request has a cookie present else aborts the request with a 403.
func (s *Server) MustCookie(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("[middleware.MustCookie] applying\n")
		cookie, err := r.Cookie(keyCookie)
		if err != nil || cookie == nil {
			logrus.Errorf("[middleware.MustCookie] no cookie present. Want: cookie - have: %v\n", cookie)
			http.Error(w, http.ErrNoCookie.Error(), http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), typeKeyCookie(keyCookie), cookie)
		next(w, r.WithContext(ctx))
	}
}

// WithTicketAuth looks for the from this service issued web-socket ticket
// and validates the ticket - if of it passed the request forward to the next caller
// else returns with a http error
func (s *Server) WithTicketAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("[middleware.WithTicketAuth] applying\n")
		ticket := r.URL.Query().Get("ticket")
		if ticket == "" {
			logrus.Errorf("could not find any ws-ticket want: jwt - have: %s\n", ticket)
			s.onErr(w, http.StatusUnauthorized, "no ws ticket found")
			return
		}
		if err := jwts.Validate(ticket); err != nil {
			logrus.Error("provided ws-ticket is not valid")
			s.onErr(w, http.StatusUnauthorized, "not authorized")
			return
		}
		ctx := context.WithValue(r.Context(), typeKeyTicket(keyTicket), ticket)
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