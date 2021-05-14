package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api-gateway/domain"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) UpdateAppConfig(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.UpdateAppConfig] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	// flag describes which part of the configuration should be updated
	flag := r.URL.Query().Get("flag")

	var form domain.UpdateConfigForm
	if err := handler.decode(r.Body, &form); err != nil {
		logrus.Errorf("<%v>[handler.UpdateAppConfig] could not decode r.Body: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		handler.onError(w, "Could not decode r.Body", http.StatusBadRequest)
		return
	}

	err := handler.domain.UpdateAppConfig(r.Context(), form, flag)
	if err != nil {
		logrus.Errorf("<%v>[handler.UpdateAppConfig] could not update config: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	handler.onSuccessJSON(w, map[string]string{"msg": "App Config Updated"}, http.StatusOK)
}
