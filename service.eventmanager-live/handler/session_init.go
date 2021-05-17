package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.eventmanager-live/domain/types"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) HandlerInitSession(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[service.eventmanager-live.InitSession] received request\n", r.Host)
	var session types.SessionStart
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		logrus.Errorf("<%v>[service.eventmanager-live.InitSession] could not decode r.Body: %v\n", r.Host, err)
		handler.onErr(w, http.StatusBadRequest, "could not decode r.Body")
		return
	}

	claims, ok := r.Context().Value(typeKeyClaims(keyClaims)).(struct{ AppUuid, AppOrigin string })
	if claims.AppUuid == "" || claims.AppOrigin == "" || !ok {
		logrus.Errorf("<%v>[service.eventmanager-live.InitSession] could not find app claims\n", r.Host)
		handler.onErr(w, http.StatusInternalServerError, "missing credentials")
		return
	}
	session.AppOrigin = claims.AppOrigin
	session.AppUuid = claims.AppUuid

	cookie, ok := r.Context().Value(typeKeyCookie(keyCookie)).(*http.Cookie)
	if !ok || cookie == nil {
		logrus.Errorf("<%v>[service.eventmanager-live.InitSession] could not find cookie\n", r.Host)
		handler.onErr(w, http.StatusInternalServerError, "missing credentials")
		return
	}
	session.Cookie = cookie.Value
	fmt.Printf("Claims: %+v\nData: %+v\n", claims, session)

	config, err := handler.domain.InitSession(r.Context(), session)
	if err != nil {
		logrus.Errorf("<%v>[service.eventmanager-live.InitSession] could not init session: %v\n", r.Host, err)
		handler.onErr(w, int(err.Code()), err.Info())
		return
	}
	// fetch meta data from app service
	handler.onSuccess(w, http.StatusOK, map[string]interface{}{
		"btn_defs": config.GetBtnTime(),
	})
}
