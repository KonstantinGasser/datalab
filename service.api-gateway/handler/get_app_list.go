package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) GetAppList(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.GetAppList] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		handler.onError(w, "User not authenticated", http.StatusUnauthorized)
		return
	}
	apps, err := handler.domain.GetAppList(r.Context(), user.Uuid)
	if err != nil {
		logrus.Errorf("<%v>[handler.GetAppList] %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	handler.onSuccessJSON(w, map[string]interface{}{
		"status": http.StatusOK,
		"apps":   apps,
	}, http.StatusOK)
}
