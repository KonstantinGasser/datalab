# lib: binder

`binder` is a simple library allowing to check if struct fields are set.
This can be helpful if your API pareses the r.Body into a struct and you want to avoid redundant field checks.

## bind tag
in your structs include the tag \``bind:"yes"`\` for each field which must bind to NOT the default value. The tag will verify that the value of the field is NOT the default value: <br>
- int => !0
- string => !""
- interface{} => !nil
- slice => len(slice) >= 0


## usage
```go
go get github.com/KonstantinGasser/datalabs/binder
```

```go
type User struct {
    UserName string `json:"username" bind:"yes"`
    Email string `json:"email" bind:"yes"`
}

func HandleNewUser(w http.ResponseWriter, r *http.Request) {
    // decode r.Body from request to user struct
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        log.Fatal(err)
    }
    
    // check that all fields were passed by the request
    // if either UserName or Email in the User struct are missing (values are default values) bind fails
    if err := binder.MustBind(&user); err != nil {
        log.Fatal(err)
    }
}

```
