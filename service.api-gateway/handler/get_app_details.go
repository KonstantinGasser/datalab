package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) GetAppDetails(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.GetAppDetails] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)
	user := ctx_value.GetAuthedUser(r.Context())
	if user == nil {
		handler.onError(w, "User not authenticated", http.StatusUnauthorized)
		return
	}
	appUuid := r.URL.Query().Get("app")
	if appUuid == "" {
		handler.onError(w, "App Uuid missing in query", http.StatusBadRequest)
		return
	}
	app, err := handler.domain.GetAppDetails(r.Context(), user.Uuid, appUuid)
	if err != nil {
		logrus.Errorf("<%v>[handler.GetAppDetails] could not get app: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	handler.onSuccessJSON(w, map[string]interface{}{"status": http.StatusOK, "app": app}, http.StatusOK)
}
