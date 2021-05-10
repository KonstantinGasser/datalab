package tokenissuer

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Todo: must go in env
const (
	issuerApp = "com.datalab.token-service"
	secretApp = "super_secure"
)

// IssueApp takes in arguments for the token of the user
// returning a JWT with exp set to expTime and the data passed in
func IssueNew(appUuid, orgnAndAppHash, appOrigin string, exp time.Time) (string, error) {
	claims := jwt.MapClaims{
		"sub":    appUuid,
		"origin": appOrigin,
		"hash":   orgnAndAppHash,
		"iss":    issuerApp,
		"iat":    time.Now().Unix(),
		"exp":    exp,
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := _token.SignedString([]byte(secretApp))
	if err != nil {
		return "", fmt.Errorf("[jwts.IssueApp] could not sign token: %v", err)
	}
	return token, nil
}
