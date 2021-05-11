package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api-gateway/domain"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) CreateApp(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.CreateApp] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var form domain.CreateAppForm
	if err := handler.decode(r.Body, &form); err != nil {
		handler.onError(w, "Could not decode r.Body", http.StatusBadRequest)
		return
	}
	user := ctx_value.GetAuthedUser(r.Context())
	uuid, err := handler.domain.CreateApp(r.Context(), user.Uuid, user.Organization, form)
	if err != nil {
		logrus.Errorf("<%v>[handler.CreateApp] could not create app: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	handler.onSuccessJSON(w, map[string]string{"msg": "App created", "uuid": uuid}, http.StatusOK)
}
