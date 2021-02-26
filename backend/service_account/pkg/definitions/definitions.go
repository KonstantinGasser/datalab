package definitions

// AccountService handles and manages  application users
type AccountService interface {
	CreateUser(CreateUserRequest) CreateUserResponse
	LoginUser(LoginUserRequest) LoginUserResponse
}

// CreateUserRequest request holding data to register new user
type CreateUserRequest struct {
	FirstName    string
	LastName     string
	Username     string
	EmailAddress string
	PasswordHash string
}

// CreateUserResponse response after registration
type CreateUserResponse struct {
	Status int
	Msg    string
}

// LoginUserRequest request to auth user
type LoginUserRequest struct {
	Username     string
	Passwordhash string
}

// LoginUserResponse response for login request
type LoginUserResponse struct {
	Status int
	Msg    string
	Token  string
}
