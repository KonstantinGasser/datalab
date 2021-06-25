package httpserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (server *Server) OpenSocket(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[server.OpenSocket] received request\n")
}
