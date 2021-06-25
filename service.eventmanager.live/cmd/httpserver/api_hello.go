package httpserver

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Hello takes care about verifing that the user has the right to be tracked (lol sounds funny)
// It therfore performs all authentication and authrorization checks, loading meta data the client needs
// and issuing an wws ticket for socket authentication
func (server *Server) Hello(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("[server.Hello] received request\n")

	fmt.Println(r.Context().Value(typeKeyIP(keyIP)))
}
