package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	apptokenissuer "github.com/KonstantinGasser/datalab/service.app-token-issuer/proto"
	"github.com/sirupsen/logrus"

	"github.com/KonstantinGasser/datalab/utils/unique"
)

// WithCookie looks-up if a request already has an x-datalab cookie set
// else sets a new x-datalab cookie. In both cases the cookie information
// gets passed into the r.Context
func (handler *Handler) WithCookie(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cookie *http.Cookie
		var err error

		cookie, err = r.Cookie(keyCookie)
		if err != nil || cookie.Value == "" {
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
			}
			http.SetCookie(w, cookie)
		}
		// pass cookie via context
		ctx := context.WithValue(r.Context(), typeKeyCookie(keyCookie), cookie)
		// move to next handler
		next(w, r.WithContext(ctx))
	}
}

// WithAuth looks-up if the request has the x-datalab-token set. If not returns a
// http.StatusUnauthorized else authenticates the token
func (handler *Handler) WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(keyDatalabToken)
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// authenticate app-token
		resp, err := handler.appTokenSvc.Validate(r.Context(), &apptokenissuer.ValidateRequest{
			Tracing_ID: "1",
			AppToken:   token,
		})
		if err != nil {
			logrus.Errorf("[middleware.WithAuth] could not validate token: %v\n", err)
			handler.onErr(w, http.StatusInternalServerError, "could not authenticate app token")
			return
		}
		if resp.GetStatusCode() != http.StatusOK {
			logrus.Errorf("[middleware.WithAuth] app-token not valid: %v\n", resp.GetMsg())
			handler.onErr(w, int(resp.GetStatusCode()), resp.GetMsg())
			return
		}
		ctx := context.WithValue(r.Context(), typeKeyClaims(keyClaims), struct {
			AppUuid, AppOrigin string
		}{
			AppUuid:   resp.GetAppUuid(),
			AppOrigin: resp.GetAppOrigin(),
		})
		// move to next handler
		next(w, r.WithContext(ctx))
	}
}

// WithCORS apply the API-CORS settings and checks if request is a pre-flight
// performed by the browser - if the request is a pre-flight WithCORS returns a
// http.StatusOK for the pre-flight request. Pre-flight requests are indicated by the
// http.Method 'OPTIONS'.
func (handler *Handler) WithCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set CORS-Header
		var reqOrigin string = r.Header.Get("Origin")
		for _, origin := range handler.allowedOrigins {
			if reqOrigin == origin {
				w.Header().Set(accessControlAllowOrigin, reqOrigin)
			}
		}
		w.Header().Set(accessControlAllowMethods, strings.Join(handler.allowedMethods, ","))
		w.Header().Set(accessControlAllowHeaders, strings.Join(handler.allowedHeaders, ","))
		w.Header().Set(accessControlAllowCreds, fmt.Sprintf("%t", handler.allowCredentials))

		// check if pre-flight request from browser
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// move to next handler
		next(w, r)
	}
}
