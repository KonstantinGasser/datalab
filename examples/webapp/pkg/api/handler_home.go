package api

import "net/http"

func (api API) HandlerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
