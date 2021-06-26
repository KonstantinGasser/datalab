package apptokens

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/KonstantinGasser/datalab/library/hasher"
	"github.com/KonstantinGasser/required"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	issuerService  = "datalab.service.app.token"
	secretAppToken = "secure"

	appTokenExpTime = 24 * time.Hour * 7
)

var (
	ErrMissingFields      = fmt.Errorf("AppToken must have appRefUuid/hash/owner")
	ErrWrongAppHash       = fmt.Errorf("provided org/app-name hash does not match with db record")
	ErrAppTokenStillValid = fmt.Errorf("current AppToken is still valid")
	ErrNoReadWriteAccess  = fmt.Errorf("user read/write access for AppToken")
	ErrNoReadAccess       = fmt.Errorf("user has no read access for AppToken")
)

type ApptokenRepo interface {
	Initialize(ctx context.Context, appToken AppToken) error
	GetById(ctx context.Context, uuid string, result interface{}) error
	Update(ctx context.Context, uuid, jwt string, exp int64, refreshCount int32) error
	SetAppTokenLock(ctx context.Context, uuid string, lock bool) error
}

// AppToken represents the token data as it will be stored in the datbase
type AppToken struct {
	AppRefUuid   string `bson:"_id" required:"yes"`
	Locked       bool   `bson:"locked"`
	AppHash      string `bson:"app_hash" required:"yes"`
	AppOwner     string `bson:"app_owner" required:"yes"`
	AppOrigin    string `bson:"app_origin"`
	Jwt          string `bson:"app_jwt"`
	Exp          int64  `bson:"app_jwt_exp"`
	RefreshCount int32  `bson:"refresh_count"`
}

// NewDefault creates a new default AppToken with only the meta data but no valid
// Jwt nor Expiration time
func NewDefault(AppRefUuid, appHash, appOwner, appOrigin string) (*AppToken, error) {
	appToken := AppToken{
		AppRefUuid:   AppRefUuid,
		AppHash:      appHash,
		AppOwner:     appOwner,
		AppOrigin:    appOrigin,
		RefreshCount: -1, // will be updated on every issue -1 therefor represents "nil"/no prior app-token exists
	}
	if err := required.Atomic(&appToken); err != nil {
		return nil, ErrMissingFields
	}
	return &appToken, nil
}

// Issue issues a new AppToken with an updated Jwt and Exp and RefreshCount. The operation fails
// if the current AppToken.Exp has not yet expired
func (appToken *AppToken) Issue(orgn, appName string) (*AppToken, error) {
	// as a verification step the user must provide
	// the org/appname which must match the stored data
	if !appToken.CompareHash(orgn, appName) {
		return nil, ErrWrongAppHash
	}
	// current AppToken must be expired in order to issue a new one
	// if non set (first time issuing) case will be ignored
	if ok := appToken.expired(); !ok && appToken.Jwt != "" {
		return nil, ErrAppTokenStillValid
	}

	issuedToken := AppToken{
		AppRefUuid:   appToken.AppRefUuid,
		Locked:       true,
		AppHash:      appToken.AppHash,
		AppOwner:     appToken.AppOwner,
		AppOrigin:    appToken.AppOrigin,
		RefreshCount: appToken.RefreshCount + 1, // increment refresh-count to invalidate current app token
		Jwt:          "",
		Exp:          0,
	}

	jwt, exp, err := issuedToken.JWT()
	if err != nil {
		return nil, err
	}
	issuedToken.Jwt = jwt
	issuedToken.Exp = exp
	return &issuedToken, nil

}

// JWT creates a new JSON-Web-Token based on the current AppToken information
func (appToken AppToken) JWT() (string, int64, error) {

	exp := time.Now().Add(appTokenExpTime)
	claims := jwt.MapClaims{
		"sub":    appToken.AppRefUuid,
		"origin": appToken.AppOrigin,
		"hash":   appToken.AppHash,
		"iss":    issuerService,
		"iat":    time.Now().Unix(),
		"exp":    exp.Unix(),
		// for some reason the JWT spec uses only three chars so it's ugly but rfc == refresh-count
		"rfc": appToken.RefreshCount,
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := _token.SignedString([]byte(secretAppToken))
	if err != nil {
		return "", 0, fmt.Errorf("[jwts.IssueApp] could not sign token: %v", err)
	}
	return token, exp.Unix(), nil
}

// IsValid matches the JWT refresh count with the refresh count from the database
// if they dont match, the app-token is marked as invalid
func (appToken AppToken) IsValid(jwtRefreshCount int) bool {
	return jwtRefreshCount == int(appToken.RefreshCount)
}

// TODO: needs to check refresh count
// IsCorrupted checks if the Json-Web-Token it self is valid and be decoded using the secure
// further more it checks if the token has been expired
func IsCorrupted(jwtString string) (string, string, int, error) {
	token, err := verifyToken(jwtString, secretAppToken)
	if err != nil {
		return "", "", 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", "", 0, errors.New("apptoken invalid")
	}
	return claims["sub"].(string), claims["origin"].(string), claims["rfc"].(int), nil
}

func ClaimsFromJwt(jwtString string) (string, string, int, error) {
	token, err := verifyToken(jwtString, secretAppToken)
	if err != nil {
		return "", "", 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", "", 0, errors.New("apptoken invalid")
	}
	return claims["sub"].(string), claims["origin"].(string), int(claims["rfc"].(float64)), nil
}

// MarkDirty updates the refresh count of an app token invalidating the all other app tokens
// it will unset the jwt, ext
func (appToken AppToken) MarkDirty() *AppToken {
	return &AppToken{
		AppRefUuid:   appToken.AppRefUuid,
		AppHash:      appToken.AppHash,
		AppOwner:     appToken.AppOwner,
		AppOrigin:    appToken.AppOrigin,
		Locked:       appToken.Locked,
		Jwt:          "",
		Exp:          0,
		RefreshCount: appToken.RefreshCount + 1,
	}
}

// CompareHash compares if the provided meta data (orgnanization name and app name)
// match with the apptoken.Hash.
func (appToken AppToken) CompareHash(orgn, appName string) bool {
	hash := hasher.Build(appName, orgn)
	return hash == appToken.AppHash
}

// HasReadWrite checks if the provided user uuid is listed as owner of
// AppToken
func (appToken AppToken) HasReadWrite(userUuid string) error {
	if appToken.AppOwner != userUuid {
		return ErrNoReadWriteAccess
	}
	return nil
}

// HasRead checks if the user has read access on the AppToken
func (appToken AppToken) HasRead(readWriteUuids ...string) error {
	for _, uuid := range readWriteUuids {
		if uuid == appToken.AppRefUuid {
			return nil
		}
	}
	return ErrNoReadAccess
}

// HasReadOrWrite checks if the user has either read or write acces on the AppToken
func (appToken AppToken) HasReadOrWrite(userUuid string, readWriteUuids ...string) error {
	rErr := appToken.HasRead(readWriteUuids...)

	rwErr := appToken.HasReadWrite(userUuid)
	if rErr != nil && rwErr != nil {
		return ErrNoReadAccess
	}
	return nil
}

// expired checks if the current jwt is already expired or not
func (appToken *AppToken) expired() bool {
	return time.Now().Unix() >= appToken.Exp
}

func verifyToken(tokenString, secret string) (*jwt.Token, error) {
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
