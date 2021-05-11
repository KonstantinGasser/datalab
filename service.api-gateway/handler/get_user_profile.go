package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.GetUserProfile] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	caller := ctx_value.GetAuthedUser(r.Context())
	user, err := handler.domain.GetUserProfile(r.Context(), caller.Uuid)
	if err != nil {
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	handler.onSuccessJSON(w, map[string]interface{}{
		"status": http.StatusOK,
		"user":   user,
	}, http.StatusOK)
}
