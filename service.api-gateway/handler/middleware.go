package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/hasher"
	"github.com/KonstantinGasser/datalab/service.api-gateway/domain"
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
		// ctx := context.WithTimeout(r.Context(), authTimeout)
		// defer cancel()
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

func (handler *Handler) WithAppPermissions(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user = ctx_value.GetAuthedUser(r.Context())
		var payload struct {
			AppUuid   string `json:"app_uuid" required:"yes"`
			AppName   string `json:"app_name" required:"yes"`
			Orgn      string `json:"owner_domain" required:"yes"`
			AppOrigin string `json:"app_origin" required:"yes"`
		}
		if err := handler.decode(r.Body, &payload); err != nil {
			logrus.Errorf("<%v>[handler.WithAppPermissions] could not decode r.Body: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
			handler.onError(w, "could not decode r.Body", http.StatusBadRequest)
		}

		err := handler.domain.HasAppPermissions(r.Context(), user.GetUuid(), payload.AppUuid, payload.AppName, payload.Orgn)
		if err != nil {
			logrus.Errorf("<%v>[handler.WithAppPermissions] could not authorize request: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
			handler.onError(w, err.Info(), int(err.Code()))
			return
		}
		// TODO don't like this: sharing of data should be done differently!
		ctx := context.WithValue(r.Context(), "app.meta", domain.AppMetaData{
			Uuid:   payload.AppUuid,
			Origin: payload.AppOrigin,
			Hash:   hasher.Build(payload.AppName, payload.Orgn),
		})
		// map[string]string{
		// 	"appOrigin": payload.AppOrigin,
		// 	"appUuid":   payload.AppUuid,
		// 	"appHash":   hasher.Build(payload.AppName, payload.Orgn),
		// }
		// )
		// serve next request
		next(w, r.WithContext(ctx))
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
