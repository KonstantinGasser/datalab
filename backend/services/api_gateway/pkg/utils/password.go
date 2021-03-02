package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 14
)

// HashFromPassword hashes the user's password before routed to other services (user-service)
func HashFromPassword(password []byte) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %v", err)
	}
	return string(bytes), nil
}
