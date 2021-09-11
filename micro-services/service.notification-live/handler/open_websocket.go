package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/library/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) HandlerOpenSocket(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.OpenSocket] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	err := handler.domain.OpenSocket(r.Context(), w, r)
	if err != nil {
		logrus.Infof("<%v>[handler.OpenSocket] could not establish connection: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err)
		handler.onError(w, "could not establish connection", http.StatusBadRequest)
		return
	}
}
