package user

import (
	"strconv"
	"strings"
	"time"

	"github.com/KonstantinGasser/required"
	"github.com/gofrs/uuid"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type User struct {
	Uuid         uuid.UUID `bson:"_id"`
	Username     string    `bson:"username" required:"yes" min:"5"`
	FirstName    string    `bson:"first_name" required:"yes"`
	LastName     string    `bson:"last_name" required:"yes"`
	Organization string    `bson:"organization" required:"yes"`
	Position     string    `bson:"position" required:"yes"`
	Avatar       string    `bson:"avatar"`
}

var (
	// defautlAvatarApi allows to generare a random avatar based on some seed (like a timestamp etc)
	defaultAvatarApi = "https://avatars.dicebear.com/api/bottts/"
)

// UpdatableUser defins the fields of a User that can be changed
type UpdatableUser struct {
	Uuid                          string
	FirstName, LastName, Position string
}

func NewDefault(username, firstname, lastname, organization, position string) (User, error) {
	// build random user avatar
	timeNow := strconv.FormatInt(time.Now().Unix(), 10)
	urlSeed := strings.Join([]string{timeNow, "svg"}, ".") // timetamp.svg
	avatarUrl := strings.Join([]string{defaultAvatarApi, urlSeed}, "/")
	uuid, err := uuid.NewV4()
	if err != nil {
		return User{}, errors.Wrap(err, "unable to create UUID")
	}

	user := User{
		Uuid:         uuid,
		Username:     username,
		FirstName:    firstname,
		LastName:     lastname,
		Organization: organization,
		Position:     position,
		Avatar:       avatarUrl,
	}
	if err := required.Atomic(&user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (u *User) NewPosition(pos string) {
	u.Position = pos
}
