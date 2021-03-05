package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HandlerAppCreate is the api endpoint if a logged in user wants to create a new application
func (api API) HandlerAppCreate(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[api.HandlerAppCreate] received create-app request: %v\n", r.Host)
	w.Write([]byte("hello friend"))
}
