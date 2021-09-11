package users

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/KonstantinGasser/datalab/library/utils/hash"
	"github.com/KonstantinGasser/datalab/library/utils/unique"
	"github.com/KonstantinGasser/required"
	"github.com/dgrijalva/jwt-go"
)

const (
	issuerUser  = "datalab.service.user.auth.agent"
	secretUser  = "super_secure"
	expTimeUser = time.Minute * 60 * 4
)

var (
	ErrMissingFields    = fmt.Errorf("not all fields are provided")
	ErrInvalidOrgnName  = fmt.Errorf("organization name does not follow policies")
	ErrWrongCredentials = fmt.Errorf("user credentials are not matching with db record")

	ErrInvalidJWT = fmt.Errorf("jwt is no longer valid")
	ErrJWTParse   = fmt.Errorf("could not parse jwt token")
	ErrCorruptJWT = fmt.Errorf("jwt could not be parsed (JWT might be corrupted)")
	ErrExpiredJWT = fmt.Errorf("provided JWT has expired")
)

type UserRepo interface {
	UsernameTaken(ctx context.Context, username string) (bool, error)
	Store(ctx context.Context, user User) error
	GetByUsername(ctx context.Context, username string, stored interface{}) error
}

type User struct {
	Uuid         string `bson:"_id"`
	Username     string `bson:"username" required:"yes"`
	Organization string `bson:"organization" required:"yes"`
	Password     string `bson:"password"`
}

type AuthedUser struct {
	Uuid         string
	Organization string
	Username     string
}

func NewDefaultUser(username, organization string) (*User, error) {
	user := User{
		Username:     strings.TrimSpace(username),
		Organization: strings.TrimSpace(organization),
	}
	if err := required.Atomic(&user); err != nil {
		return nil, ErrMissingFields
	}
	if ok := allowedOrganizationName(user.Organization); !ok {
		return nil, ErrInvalidOrgnName
	}
	err := user.UUID()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UUID assigns a new unique identifier to the User struct
func (u *User) UUID() error {
	uuid, err := unique.UUID()
	if err != nil {
		return err
	}
	u.Uuid = uuid
	return nil
}

// HashAndSalt hashes the plain text user password
func (u *User) HashAndSalt(password string) error {
	hashed, err := hash.FromPassword([]byte(strings.TrimSpace(password)))
	if err != nil {
		return err
	}
	u.Password = hashed
	return nil
}

func (u User) Credentials(password string) error {
	if !hash.CheckPasswordHash(password, u.Password) {
		return ErrWrongCredentials
	}
	return nil
}

// AccessToken generates a JWT based on the User information
func (u User) AccessToken() (string, error) {
	claims := jwt.MapClaims{
		"sub":   u.Uuid,
		"uname": u.Username,
		"orgn":  u.Organization,
		"iat":   issuerUser,
		"exp":   time.Now().Add(expTimeUser).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(secretUser))
	if err != nil {
		return "", fmt.Errorf("[jwts.IssueUser] could not sign token: %v", err)
	}

	return accessToken, nil
}

func LoggedIn(accessToken string) (*AuthedUser, error) {
	token, err := verifyToken(accessToken, secretUser)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, ErrInvalidJWT
	}
	if err := claims.Valid(); err != nil {
		return nil, ErrExpiredJWT
	}

	return &AuthedUser{
		Uuid:         claims["sub"].(string),
		Organization: claims["orgn"].(string),
		Username:     claims["uname"].(string),
	}, nil
}

func verifyToken(tokenString, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrCorruptJWT
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, ErrJWTParse
	}
	return token, nil
}

// allowedOrganizationName verifies that the provided organization
// name matches with the specifications
func allowedOrganizationName(name string) bool {
	re := regexp.MustCompile("/")
	matches := re.Find([]byte(name))
	return matches == nil
}
