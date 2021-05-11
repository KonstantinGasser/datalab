package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api-gateway/domain"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.RegisterUser] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var form domain.RegisterForm
	if err := handler.decode(r.Body, &form); err != nil {
		handler.onError(w, "Could not decode r.Body", http.StatusBadRequest)
		return
	}
	err := handler.domain.RegisterUser(r.Context(), form)
	if err != nil {
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	handler.onSuccessJSON(w, map[string]string{"msg": "User-Account created"}, http.StatusOK)
}
