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
