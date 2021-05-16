package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler Handler) GetUserColleagues(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.GetUserColleagues] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	user := ctx_value.GetAuthedUser(r.Context())
	colleagues, err := handler.domain.GetColleagues(r.Context(), user.Uuid)
	if err != nil {
		logrus.Errorf("<%v>[handler.GetUserColleagues] could not colleagues list: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	handler.onSuccessJSON(w, map[string]interface{}{
		"status":     http.StatusOK,
		"colleagues": colleagues,
	}, http.StatusOK)
}
