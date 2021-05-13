package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) CreateAppToken(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.CreateAppToken] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var user = ctx_value.GetAuthedUser(r.Context())
	// TODO: pass appMeta info from middleware permissions in a cleaner form
	var appMeta = r.Context().Value("app.meta").(map[string]string)

	token, err := handler.domain.CreateAppToken(r.Context(), user.GetUuid(), appMeta["appUuid"], appMeta["appOrigin"], appMeta["appHash"])
	if err != nil {
		logrus.Errorf("<%v>[handler.CreateAppToken] could not create app token: %v", ctx_value.GetString(r.Context(), "tracingID"), err)
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	handler.onSuccessJSON(w, map[string]interface{}{
		"status":    http.StatusOK,
		"app_token": token,
	}, http.StatusOK)
}
