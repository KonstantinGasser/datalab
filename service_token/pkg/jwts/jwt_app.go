package jwts

import (
	"context"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Todo: must go in env
const (
	issuerApp  = "com.datalab.token-service"
	SecretApp  = "super_secure"
	expTimeApp = time.Hour * 24 * 7 // valid for 7 days
)

// IssueApp takes in arguments for the token of the user
// returning a JWT with exp set to expTime and the data passed in
func IssueApp(ctx context.Context, appUuid, orgnAndAppHash, appOrigin string) (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = appUuid
	claims["origin"] = appOrigin
	claims["hash"] = orgnAndAppHash
	claims["iss"] = issuerApp
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(expTimeApp).Unix()

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := _token.SignedString([]byte(SecretApp))
	if err != nil {
		return "", fmt.Errorf("[jwts.IssueApp] could not sign token: %v", err)
	}
	return token, nil
}
