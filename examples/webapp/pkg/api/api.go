package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/services/user"
	"github.com/KonstantinGasser/clickstream/examples/webapp/pkg/storage"
)

// API is the struct which defines the behaviors of
// the web-server. Including routing, middleware and helper funcs
type API struct {
	route     func(path string, handler http.HandlerFunc)
	onError   func(w http.ResponseWriter, code int, err error)
	onSuccess func(w http.ResponseWriter, code int, data interface{})

	// *** Service Dependencies ***
	userService user.User
	// *** Storage Dependency ***
	storage storage.Storage
}

func New(userService user.User, storage storage.Storage) API {
	return API{
		route: func(path string, handler http.HandlerFunc) {
			log.Printf("[route-%s] mapped to web-server\n", path)
			http.HandleFunc(path, handler)
		},
		onError: func(w http.ResponseWriter, code int, err error) {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
		},
		onSuccess: func(w http.ResponseWriter, code int, data interface{}) {
			if err := json.NewEncoder(w).Encode(data); err != nil {
				log.Printf("could not encode data: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
		// *** Service Dependencies ***
		userService: userService,
		// *** Storage Dependency ***
		storage: storage,
	}
}

// decode takes a io.Reader (can be request body) and decodes the data
// into the passed data interface{}. "data" must be a pointer to the type else
// operation will fail
func (api API) decode(body io.Reader, data interface{}) error {
	if err := json.NewDecoder(body).Decode(data); err != nil {
		log.Printf("could not decode body: %v", err)
		return err
	}
	return nil
}

// SetUp adds all routes with middleware and configurations to the API-Server
// (initializes the API). This is a responsibility of the client how's creating
// a new API
func (api API) SetUp() {
	api.route("/", api.HandlerHome)
	api.route("/register", api.HandlerRegister)
	api.route("/login", api.HandlerLogin)
}
