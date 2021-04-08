package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// API represents the handler functions and middleware
type API struct {
	// Route is a mapper between requested URL and handler
	// allows to add middleware in a nice chained way
	route func(path string, h http.Handler)
}

// New create and returns a new API struct
func New() API {
	return API{
		// route is a custom function allowing to set path and request handler
		// for a given route (similar to the http.HandlerFunc). However having it
		// customs allows to do middleware in a nicer way
		route: func(path string, h http.Handler) {
			logrus.Infof("[set-up:route] %s\n", path)
			hf := func(w http.ResponseWriter, r *http.Request) {
				if r.Method == "OPTIONS" {
					w.WriteHeader(http.StatusOK)
					return
				}
				w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
				w.Header().Set("cache-control", "no-cache")
				h.ServeHTTP(w, r)
			}
			http.HandleFunc(path, hf)
		},
	}
}

// SetUp sets up all the routes the API has along with all the middleware
// each request required to have
func (api API) SetUp() {
	logrus.Infof("\n*** adding routes to API-Service ***\n")

	rootFS := http.FileServer(http.Dir("./resources"))
	api.route("/", rootFS)
}
