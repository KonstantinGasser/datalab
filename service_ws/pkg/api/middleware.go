package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/KonstantinGasser/datalab/utils/unique"
)

// WithCookie looks-up if a request already has an x-datalab cookie set
// else sets a new x-datalab cookie. In both cases the cookie information
// gets passed into the r.Context
func (api *API) WithCookie(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(keyCookie)
		if err != nil || cookie.Value == "" {
			// set new cookie for request
			uuid, err := unique.UUID()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			c := http.Cookie{
				Name:    keyCookie,
				Value:   uuid,
				Expires: time.Now().Add(1 * time.Hour),
				Path:    "/",
			}
			http.SetCookie(w, &c)
		}
		// pass cookie via context
		ctx := context.WithValue(r.Context(), typeKeyCookie(keyCookie), cookie)
		// move to next handler
		next(w, r.WithContext(ctx))
	}
}

// WithAuth looks-up if the request has the x-datalab-token set. If not returns a
// http.StatusUnauthorized else authenticates the token
func (api *API) WithAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(keyDatalabToken)
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// TODO: - perform authentication via Token-Service
		//       - parse required values (meta data for the client)

		// move to next handler
		next(w, r)
	}
}

// WithCORS apply the API-CORS settings and checks if request is a pre-flight
// performed by the browser - if the request is a pre-flight WithCORS returns a
// http.StatusOK for the pre-flight request. Pre-flight requests are indicated by the
// http.Method 'OPTIONS'.
func (api *API) WithCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set CORS-Header
		var reqOrigin string = r.Header.Get("Origin")
		fmt.Println("Origin: ", reqOrigin)
		for _, origin := range api.allowedOrigins {
			if reqOrigin == origin {
				w.Header().Set(accessControlAllowOrigin, reqOrigin)
			}
		}
		w.Header().Set(accessControlAllowMethods, strings.Join(api.allowedMethods, ","))
		w.Header().Set(accessControlAllowHeaders, strings.Join(api.allowedHeaders, ","))
		w.Header().Set(accessControlAllowCreds, fmt.Sprintf("%t", api.allowCredentials))

		// check if pre-flight request from browser
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// move to next handler
		next(w, r)
	}
}
