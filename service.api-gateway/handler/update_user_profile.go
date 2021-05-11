package handler

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/service.api-gateway/domain"
	"github.com/KonstantinGasser/datalab/utils/ctx_value"
	"github.com/sirupsen/logrus"
)

func (handler *Handler) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("<%v>[handler.UpdateUserProfile] received request: %v\n", ctx_value.GetString(r.Context(), "tracingID"), r.Host)

	var form domain.UserUpdateForm
	if err := handler.decode(r.Body, &form); err != nil {
		handler.onError(w, "Could not decode r.Body", http.StatusBadRequest)
		return
	}

	user := ctx_value.GetAuthedUser(r.Context())
	err := handler.domain.UpdateUserProfile(r.Context(), user.Uuid, form)
	if err != nil {
		logrus.Infof("<%v>[handler.UpdateUserProfile] could not update profile: %v\n", ctx_value.GetString(r.Context(), "tracingID"), err.Error())
		handler.onError(w, err.Info(), int(err.Code()))
		return
	}
	handler.onSuccessJSON(w, map[string]string{"msg": "User-Profile updated"}, http.StatusOK)
}
