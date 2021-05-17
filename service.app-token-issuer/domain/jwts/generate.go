package jwts

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	issuerService  = "com.datalab.token-service"
	secretAppToken = "secure"
)

var (
	ErrInvalidJWT = fmt.Errorf("jwt is no longer valid")
	ErrJWTParse   = fmt.Errorf("could not parse jwt token")
	ErrCorruptJWT = fmt.Errorf("jwt could not be parsed (JWT might be corrupted)")
)

type AppClaims struct {
	AppUuid, AppOrigin string
}

func Generate(appUuid, origin, appHash string, exp time.Time) (string, error) {
	claims := jwt.MapClaims{
		"sub":    appUuid,
		"origin": origin,
		"hash":   appHash,
		"iss":    issuerService,
		"iat":    time.Now().Unix(),
		"exp":    exp,
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := _token.SignedString([]byte(secretAppToken))
	if err != nil {
		return "", fmt.Errorf("[jwts.IssueApp] could not sign token: %v", err)
	}
	return token, nil

}

func Claims(tokenString string) (*AppClaims, error) {
	token, err := verifyToken(tokenString, secretAppToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, ErrInvalidJWT
	}
	return &AppClaims{
		AppUuid:   claims["sub"].(string),
		AppOrigin: claims["origin"].(string),
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
