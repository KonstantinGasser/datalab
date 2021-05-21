package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/hasher"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) CreateAppToken(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.CreateAppToken] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var user = ctx_value.GetAuthedUser(r.Context())
	// TODO: pass appMeta info from middleware permissions in a cleaner form
	var payload struct {
		AppUuid   string `json:"app_uuid" required:"yes"`
		AppName   string `json:"app_name" required:"yes"`
		Orgn      string `json:"owner_domain" required:"yes"`
		AppOrigin string `json:"app_origin" required:"yes"`
	}
	if err := handler.decode(r.Body, &payload); err != nil {
		logrus.Errorf("<%v>[handler.WithAppPermissions] could not decode r.Body: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
		handler.onError(w, "could not decode r.Body", http.StatusBadRequest)
		return
	}

	token, err := handler.domain.CreateAppToken(r.Context(),
		user.GetUuid(),
		payload.AppUuid,
		payload.AppOrigin,
		hasher.Build(payload.AppName, payload.Orgn),
	)
	if err != nil {
		logrus.Errorf("<%v>[handler.CreateAppToken] could not create app token: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	handler.onSuccessJSON(w, map[string]interface{}{
		"status":    http.StatusOK,
		"app_token": token,
	}, http.StatusOK)
}
