package apptokens

import (
	"context"
	"fmt"
	"time"

	"github.com/KonstantinGasser/required"
	"github.com/dgrijalva/jwt-go"
)

const (
	issuerService  = "datalab.service.app.token"
	secretAppToken = "secure"

	appTokenExpTime = 24 * time.Hour * 7
)

var (
	ErrMissingFields      = fmt.Errorf("AppToken must have appRefUuid/hash/owner")
	ErrAppTokenStillValid = fmt.Errorf("current AppToken is still valid")
	ErrNoReadWriteAccess  = fmt.Errorf("user read/write access for AppToken")
	ErrNoReadAccess       = fmt.Errorf("user has no read access for AppToken")
)

type ApptokenRepo interface {
	Initialize(ctx context.Context, appToken AppToken) error
	GetById(ctx context.Context, uuid string, result interface{}) error
	Update(ctx context.Context, uuid, jwt string, exp int64) error
}

// AppToken represents the token data as it will be stored in the datbase
type AppToken struct {
	AppRefUuid string `bson:"_id" required:"yes"`
	AppHash    string `bson:"app_hash" required:"yes"`
	AppOwner   string `bson:"app_owner" required:"yes"`
	AppOrigin  string `bson:"app_origin"`
	Jwt        string `bson:"app_jwt"`
	Exp        int64  `bson:"app_jwt_exp"`
}

// NewDefault creates a new default AppToken with only the meta data but no valid
// Jwt nor Expiration time
func NewDefault(AppRefUuid, appHash, appOwner string) (*AppToken, error) {
	appToken := AppToken{
		AppRefUuid: AppRefUuid,
		AppHash:    appHash,
		AppOwner:   appOwner,
	}
	if err := required.Atomic(&appToken); err != nil {
		return nil, ErrMissingFields
	}
	return &appToken, nil
}

// Issue issues a new AppToken with an updated Jwt and Exp. The operation fails
// if the current AppToken.Exp has not yet expired
func (appToken *AppToken) Issue() (*AppToken, error) {
	// current AppToken must be expired in order to issue a new one
	// if non set (first time issuing) case will be ignored
	if ok := appToken.expired(); !ok && appToken.Jwt != "" {
		return nil, ErrAppTokenStillValid
	}
	jwt, exp, err := appToken.JWT()
	if err != nil {
		return nil, err
	}
	return &AppToken{
		AppRefUuid: appToken.AppRefUuid,
		AppHash:    appToken.AppHash,
		AppOwner:   appToken.AppOwner,
		AppOrigin:  appToken.AppOrigin,
		Jwt:        jwt,
		Exp:        exp,
	}, nil
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
	}

	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := _token.SignedString([]byte(secretAppToken))
	if err != nil {
		return "", 0, fmt.Errorf("[jwts.IssueApp] could not sign token: %v", err)
	}
	return token, exp.Unix(), nil
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
