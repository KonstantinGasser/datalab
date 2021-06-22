package httpserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (server *Server) OpenSocket(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[server.OpenSocket] received request\n")
	err := server.stream.HttpUpgrade(w, r)
	if err != nil {
		logrus.Errorf("could not upgrade to ws: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
