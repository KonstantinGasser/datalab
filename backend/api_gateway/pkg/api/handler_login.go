package api

import "net/http"

// HandlerLogin forwards the login request to the user and token service
// in oder to check the user's auth and to issue a JWT.
func (api API) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	// just for testing
	w.WriteHeader(http.StatusOK)
	return
}
