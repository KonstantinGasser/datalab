package middelware

import (
	"context"
	"log"
	"net/http"

	"github.com/KonstantinGasser/bugsbunny/service_a/pkg/helper"
)

func WithLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func WithIP(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), helper.CtxReqIP, r.RemoteAddr)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
