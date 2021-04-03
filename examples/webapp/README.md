

## Code Structure
The project can be structured in a hexagon fashion which increases the de-coupling of the components. A examples folder structure can look as followed:
- WebApp
- - cmd
- - - server
- - - - server.go
- - pkg
- - - api
- - - - api.go
- - - - handler_register.go
- - - - handler_login.go
- - - services
- - - - user 
- - - - - user.go
- - - storage
- - - - storage.go
- - main.go

The interesting part here is the `pkg` directory in which the different components are described. There is a dedicated `api` folder in which everything regarding the API will be and there are `services/user` and `storage` folders again only dealing with their domain. (this is not a must an surely can be done in 100 different ways but that's how I have been doing it and some others :).

## Starting with the `main.go`
The `main.go` file is the entry point with the main function for the program. Since the `main func` in Go does not return any error or status code like in C it is nice to only deal with the setup of the basics and call to a `main abstraction`
``` golang
func main() {
	address := flag.String("addr", "localhost:8080", "address the web-server can receive requests")
	log.Fatal(server.Run(*address))
}
```
As you can see the main function only deals with parsing customs CLI arguments like the address the web-server should listen on. After that it calls out the the `abstraction` to actually run the server.

## `main func abstraction`
This is the part were we start dealing with all the things we need to do in order to start the web-server.
- dealing with dependency injection
- creating an API
- initializing the API
- and so one

In the file `cmd/server/server.go` we have this function
```golang
func Run(serverAddress string) error {
	// create network listener listening on TCP:somePort
	listener, err := net.Listen("tcp", serverAddress)
    defer listener.Close()
	if err != nil {
		log.Fatalf("could not start listener on: %s:%v", serverAddress, err)
		return err
	}
	// create new user-service as api dep
	userSrv := user.New()
	// create new storage as api dep
	storage := storage.New("in-memory")
	// create API instance
	api := api.New(userSrv, storage)
	// setup routes and init API
	api.SetUp()

	log.Printf("starting HTTP-Server on: %s\n", serverAddress)
	if err := http.Serve(listener, nil); err != nil {
		log.Fatalf("could not start http Server: %v", err)
		return err
	}
	return nil
}
```
First we create a network listener and advise to use "TCP" and the custom server address. Next we create all of the dependencies the API needs - here it is the `UserService` and the in-memory `Storage`. Lastly we start the server by saying
`http.Serve(listener, nil)`.

## The API
As for the API, it is a struct which defines the API behavior. Lets have a look at the examples API
``` golang
type API struct {
	route     func(path string, handler http.HandlerFunc)
	onError   func(w http.ResponseWriter, code int, err error)
	onSuccess func(w http.ResponseWriter, code int, data interface{})

	// *** Service Dependencies ***
	userService user.User
	// *** Storage Dependency ***
	storage storage.Storage
}
```
Here `route` is a  custom function which adds an API-Route to the Web-Server - how that works will be shown below.
The next two functions are useful in order to not repeat code over and over again. `onError` as the name suggest is used when the response must be an error and likewise the `onSuccess` can be used to return a positive response to the client.
Both the function accept a `http.ResponseWriter` which is nothing else but the network connection to communicate with the client. Further the functions need a status code which will be returned and for the `onSuccess` function also some data that could be JSON or anything else which needs to be returned to the client.

Lastly the API holds all its dependencies - in this case the UserService and a Storage (both are interfaces and therefore can be replaced by any struct implementing it).

### Creating a new API
``` golang
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
```
### Added routes to the API
``` golang
func (api API) SetUp() {
	api.route("/", api.HandlerHome)
	api.route("/register", api.HandlerRegister)
	api.route("/login", api.HandlerLogin)
}
```
Routes can be added by calling the `api.route` function assigning a route and passing in a `http.HandlerFunc` which will handle the incoming request. Every function which has the method signature `HandlerFunc(w http.ResponseWriter, r *http.Request)` may serve as a `http.HandlerFunc` and can be used as value in `api.route`. As you can see the `SetUp` func starts with a capital `S` and will be accessible from outside the `package api` - it will be called in the `cmd/server/sever.go`.

## Storage dependency
A good practice for creating a storage dependency is to first create a `interface` for it which allows to use different storages as required.
``` golang 
type Storage interface {
	Put(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
	Exists(ctx context.Context, key string) bool
}
```
In our example the `Storage interface` looks like this. Here we define how the client can interact with the storage. Pretty straight forward - this storage only allows to either `Put` something in a storage, `Get` it back and check if an item `Exists` in the storage. <br>
By defining this interface we can now build our storage. This could be a mongoDB which we use to save and get things or like in this example just a `in-memory` Key:Value storage.

## In-Memory Storage
First we define a struct which represents the in-memory storage
``` golang
type inMem struct {
	sync.Mutex
	// store is a map holding key:value pairs
	store map[string]interface{}
}
```
The field `store` is the actual `map` which holds the data (maps in Go can be compared to HashMaps in Java, dictionaries in Python or Objects in JavaScript). For now the `sync.Mutex` is something we can ignore, but what it does it it ensures, that a Storage operation will block the access to the `store` and therefore makes it thread-save.
Now that a have struct representing our storage we can add all the function in order to implement the `Storage interface`.
I will not really talk about the content of the function since this is only an example and surely will differ depending on how you want to store data (a reference for a MongoDB storage can be found here: https://github.com/KonstantinGasser/clickstream/blob/main/backend/services/user_service/pkg/storage/mongo.go). Once all functions as defined by the `Storage interface` are implemented the `type inMem` can now act as a `type Storage`.

## User-Service
The same way we did for the `Storage` we do for the `UserService` dependency.
``` golang
type User interface {
	Register(ctx context.Context, storage storage.Storage, in RegisterRequest) (int, error)
	Login(ctx context.Context, storage storage.Storage, in LoginRequest) (int, error)
}
```
Again we need a struct which implements all of the functions and then can act as a `type User`
One interesting thing I want to point out is that the `User` function require to have a `storage.Storage` passed as an argument. This is referring to the `hexagonal` project-structure. The user service is not coupled to a specific Storage but rather the Storage can switch at runtime.
As a last thing the `package user` should also define what data the functions accept. The `Register` function requires a `RegisterRequest` type.
``` golang
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
```
Two things which are important:
- the field names must be exported -> start with a capital latter. Else the data cannot be unmarshaled
- each field must have a \`json:"json_field_name"\`. Else the unmarshal will fail to extract the data from the request JSON

Doing it this way you do not need to pass around `map[string]interface{}` to represent JSON data and will not loose the type-safety of Go itself.


## API HandlerFunc
To wrap everything up lets have a look at the `api.HandlerRegister` function
``` golang 
func (api API) HandlerRegister(w http.ResponseWriter, r *http.Request) {
	var reqData user.RegisterRequest
	if err := api.decode(r.Body, &reqData); err != nil {
		api.onError(w, http.StatusBadRequest, errors.New("could not decode request body"))
		return
	}
	// call to user-service dependency to handle register
	status, err := api.userService.Register(r.Context(), api.storage, reqData)
	if err != nil {
		api.onError(w, status, errors.New("could not register user"))
		return
	}

	api.onSuccess(w, status, map[string]string{
		"msg": "welcome new user, happy to see you!",
	})
}
```
The first thing we do is declaring a variable (`reqData`) which will hold the request JSON. Notice that `reqData` is of type `user.RegisterRequest` and hence will bind the JSON fields. Next we use the API helper func `decode` to decode the request JSON into the `reqData` - important: notice how the `api.decode` is not returning any data? That is because we are passing in the pointer to the `reqData`. If you feel uncomfortable using pointer you could re-write the function to return the data..but make sure to learn about pointer since Go uses them in quite some places :). After we decoded the JSON data we can forward the request to the user service to perform the registration. The way I have implemented the user service functions their first return parameter will always be a status-code (200, 404, 401, 500) then optional data and lastly an error. If the operation fails I handle the error and let the client know by utilizing the `api.onError` function. If everything works the `api.onSuccess` is used to inform the client.


## Example in action
Clone the repository with `git clone https://github.com@KonstantinGasser/clickstream`. CD in the `example/webapp` directory and run `go run main.go`. With `curl` you can make a request to `"http://localhost:8080/regsiter"` pass data with `-d` in the form of 
``` json 
{"username": "cool_user_name", "password": "secure_as_hell"}
```


## Further Resources
### Web-Server / API: how to
- Mat Ryer on Web-Servers: https://www.youtube.com/watch?v=rWBSMsLG8po
- Mat Ryer Blog-Post on APIs: https://medium.com/@matryer/how-i-write-go-http-services-after-seven-years-37c208122831
- Kat Zień Hexagonal Project Structure: https://www.youtube.com/watch?v=oL6JBUk6tj0, https://www.youtube.com/watch?v=vKbVrsMnhDc

### GoLang
- GoLang: https://golang.org
- PlayGround: https://play.golang.org

