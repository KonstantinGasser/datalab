package api

import (
	"net/http"
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
