package api

import (
	"log"
	"net/http"
)

// WithCors enables CORS by setting the 'Access-Control-Allow-Origin' and
// 'Access-Control-Allow-Methods' header as specified by the API struct
func (api API) WithCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Change to API props
		w.Header().Set("Access-Control-Allow-Origin", api.cfg.AccessControlAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", api.cfg.AccessControlAllowMethods)

		// return status OK if requested method is OPTIONS (used to do pre-flights from browser)
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// serve request
		next(w, r)
	}
}

// Log can be used to log information before a call is executed
func Log(next http.HandlerFunc, msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(msg)
		next(w, r)
	}
}
