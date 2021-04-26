package jwts

import (
	"context"
	"fmt"
	"time"

	tokenSrv "github.com/KonstantinGasser/datalabs/protobuf/token_service"
	jwt "github.com/dgrijalva/jwt-go"
)

// Todo: must go in env
const (
	issuerUser  = "com.datalabs.token-service"
	SecretUser  = "super_secure"
	expTimeUser = time.Minute * 60 * 4
)

// IssueUser takes in arguments for the token of the user
// returning a JWT with exp set to expTime and the data passed in
func IssueUser(ctx context.Context, user *tokenSrv.UserClaim) (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = user.GetUuid()
	claims["orgn"] = user.GetOrgnDomain()
	claims["iat"] = issuerUser
	claims["exp"] = time.Now().Add(expTimeUser).Unix()

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := _token.SignedString([]byte(SecretUser))
	if err != nil {
		return "", fmt.Errorf("[jwts.IssueUser] could not sign token: %v", err)
	}
	return token, nil
}
