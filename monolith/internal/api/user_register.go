package api

import (
	"net/http"

	"github.com/KonstantinGasser/datalab/monolith/internal/dto"
	"github.com/KonstantinGasser/datalab/monolith/pkg/ctxkey"
	"github.com/sirupsen/logrus"
)

func (api ApiServer) HandlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	var req dto.ReqRegisterUser
	if err := dto.FromRequest(r, &req); err != nil {
		logrus.Errorf("[%s] could not deserialize request: %v", r.Context().Value(ctxkey.Str("tracingID")), err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
