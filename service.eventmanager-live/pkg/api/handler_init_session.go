package api

import (
	"encoding/json"
	"net/http"

	"github.com/KonstantinGasser/datalab/service.eventmanager-live/domain/types"
)

func (api API) HandlerInitSession(w http.ResponseWriter, r *http.Request) {

	var session types.SessionStart
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		api.onErr(w, http.StatusBadRequest, "could not decode r.Body")
		return
	}

	claims := r.Context().Value(typeKeyClaims(keyClaims))
	if claims == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// fetch meta data from app service
	w.WriteHeader(http.StatusOK)
}
