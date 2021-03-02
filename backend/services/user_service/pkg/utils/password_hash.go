package utils

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 10
)

// HashFromPassword hashes the user's password before routed to other services (user-service)
func HashFromPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %v", err)
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a given password with one from the database
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func UUID() (string, error) {
	UUID, err := uuid.NewV4()
	if err != nil {
		logrus.Errorf("[userService.CreateUser] could not generate UUID for user: %v", err)
		return "", fmt.Errorf("could not generate UUID for user: %v", err)
	}
	return UUID.String(), nil
}
