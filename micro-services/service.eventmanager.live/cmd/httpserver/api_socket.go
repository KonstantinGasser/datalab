package httpserver

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (server *Server) OpenSocket(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[server.OpenSocket] received request\n")

	if err := server.eventBus.UpgradeProtocoll(w, r); err != nil {
		logrus.Errorf("[server.OpenSocket] could not open socket: %v\n", err)
		server.onErr(w, http.StatusBadRequest, "ups")
		return
	}
	// dont write back via ResponseWriter - will gets hijacked by Websocket server
}
