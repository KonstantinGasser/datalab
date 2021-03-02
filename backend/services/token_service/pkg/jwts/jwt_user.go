package jwts

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
)

// Todo: must go in env
const (
	issuer  = "com.datalabs.token-service"
	secret  = "super_secure"
	expTime = time.Minute * 60
)

// IssueUser takes in arguments for the token of the user
// returning a JWT with exp set to expTime and the data passed in
func IssueUser(uuid, username, orgnDomain string) (string, error) {
	// calims holds all the data which will be
	// encoded in the JWT
	claims := jwt.MapClaims{}
	claims["uuid"] = uuid
	claims["username"] = username
	claims["orgnDomain"] = orgnDomain
	claims["iat"] = issuer
	claims["exp"] = time.Now().Add(expTime).Unix()

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := _token.SignedString([]byte(secret))
	if err != nil {
		logrus.Errorf("[jwts.IssueUser] could not sign token: %v", err)
		return "", fmt.Errorf("[jwts.IssueUser] could not sign token: %v", err)
	}
	return token, nil
}
