package jwts

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/KonstantinGasser/clickstream/backend/services/token_service/pkg/utils"
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
func IssueUser(ctx context.Context, uuid, username, orgnDomain string) (string, error) {
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
		logrus.Errorf("<%v>[jwts.IssueUser] could not sign token: %v", utils.StringValueCtx(ctx, "tracingID"), err)
		return "", fmt.Errorf("[jwts.IssueUser] could not sign token: %v", err)
	}
	return token, nil
}

func GetJWTClaims(ctx context.Context, tokenString string) (map[string]interface{}, error) {
	token, err := verifyToken(tokenString)
	if err != nil {
		logrus.Errorf("<%v>[jwts.GetJWTClaims] %v", utils.StringValueCtx(ctx, "tracingID"), err)
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		logrus.Warnf("<%v>[jwts.ExtractTokenMetaData] token not valid or no claims found\n", utils.StringValueCtx(ctx, "tracingID"))
		return nil, errors.New("user not authenticated")
	}
	user := make(map[string]interface{})
	user["uuid"] = claims["uuid"]
	user["orgnDomain"] = claims["orgnDomain"]
	user["username"] = claims["username"]
	return user, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("[jwts.verifyToken] could not parse JWT: %v", err)
	}
	return token, nil
}
