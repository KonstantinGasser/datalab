package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

type callRes struct {
	Key   string
	Value interface{}
}

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

	app, err := handler.domain.GetAppInfo(r.Context(), user.Uuid, appUuid)
	if err != nil {
		logrus.Errorf("<%v>[handler.GetAppDetails] could not get app: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}

	result := make(chan callRes)
	go func() {
		token, err := handler.domain.GetAppToken(r.Context(), appUuid)
		if err != nil {
			logrus.Warnf("<%v>[handler.GetAppDetails] could not get app token: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		}
		result <- callRes{Key: "token", Value: token}
	}()

	go func() {
		config, err := handler.domain.GetAppConfig(r.Context(), appUuid)
		if err != nil {
			logrus.Warnf("<%v>[handler.GetAppDetails] could not get app config: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		}
		result <- callRes{Key: "config", Value: config}
	}()
	go func() {
		owner, err := handler.domain.GetUserProfile(r.Context(), app.Owner)
		if err != nil {
			logrus.Warnf("<%v>[handler.GetAppDetails] could not get app owner: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		}
		result <- callRes{Key: "owner", Value: owner}
	}()

	var appInfo = map[string]interface{}{"status": http.StatusOK, "app": app}

	res1 := <-result
	appInfo[res1.Key] = res1.Value
	res2 := <-result
	appInfo[res2.Key] = res2.Value
	res3 := <-result
	appInfo[res3.Key] = res3.Value

	handler.onSuccessJSON(w, appInfo, http.StatusOK)
}
