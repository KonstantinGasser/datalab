package api

import (
	"errors"
	"net/http"

	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/services/user"
)

func (api API) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	var reqData user.RegisterRequest
	if err := api.decode(r.Body, &reqData); err != nil {
		api.onError(w, http.StatusBadRequest, errors.New("could not decode request body"))
		return
	}
	// call to user-service dependency to handle register
	status, err := api.userService.Register(r.Context(), api.storage, reqData)
	if err != nil {
		api.onError(w, status, errors.New("could not register user"))
		return
	}

	api.onSuccess(w, status, map[string]string{
		"msg": "welcome new user, happy to see you!",
	})
}
