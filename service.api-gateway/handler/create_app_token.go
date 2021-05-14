package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api-gateway/domain"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) CreateAppToken(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.CreateAppToken] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var user = ctx_value.GetAuthedUser(r.Context())
	// TODO: pass appMeta info from middleware permissions in a cleaner form
	appMeta, ok := r.Context().Value("app.meta").(domain.AppMetaData)
	if !ok {
		logrus.Errorf("<%v>[handler.CreateAppToken] could not get app meta data from ctx\n", ctx_value.GetString(r.Context(), "tracingID"))
		handler.onError(w, "App Permission check failed", http.StatusInternalServerError)
		return
	}

	token, err := handler.domain.CreateAppToken(r.Context(), user.GetUuid(), appMeta.Uuid, appMeta.Origin, appMeta.Hash)
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
