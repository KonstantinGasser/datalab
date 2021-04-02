package api

import (
	"errors"
	"net/http"

	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/services/user"
)

func (api API) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	var in user.LoginRequest
	if err := api.decode(r.Body, &in); err != nil {
		api.onError(w, http.StatusBadRequest, errors.New("could not decode request body"))
		return
	}
	status, err := api.userService.Login(r.Context(), api.storage, in)
	if err != nil {
		api.onError(w, status, errors.New("could not login user"))
		return
	}
	api.onSuccess(w, status, map[string]interface{}{
		"logged-in": status,
	})
}
