package jwts

import (
	"fmt"
	"time"

	"github.com/KonstantinGasser/datalab/service.user-authentication/domain/types"
	jwt "github.com/dgrijalva/jwt-go"
)

// Todo: must go in env
const (
	issuerUser  = "com.datalab.service.user-authentication"
	secretUser  = "super_secure"
	expTimeUser = time.Minute * 60 * 4
)

var (
	ErrInvalidJWT = fmt.Errorf("jwt is no longer valid")
	ErrJWTParse   = fmt.Errorf("could not parse jwt token")
	ErrCorruptJWT = fmt.Errorf("jwt could not be parsed (JWT might be corrupted)")
)

// IssueUser takes in arguments for the token of the user
// returning a JWT with exp set to expTime and the data passed in
func Issue(uuid, organization string, permissions []types.AppPermission) (string, error) {
	claims := jwt.MapClaims{
		"sub":  uuid,
		"orgn": organization,
		"iat":  issuerUser,
		"exp":  time.Now().Add(expTimeUser).Unix(),
		"apps": permissions,
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := _token.SignedString([]byte(secretUser))
	if err != nil {
		return "", fmt.Errorf("[jwts.IssueUser] could not sign token: %v", err)
	}
	return token, nil
}

func GetJWTClaims(tokenString string) (map[string]interface{}, error) {
	token, err := verifyToken(tokenString, secretUser)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, ErrInvalidJWT
	}
	return claims, nil
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
