package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func createHashedPW(password string) (string, error) {
	if password == "" {
		return "", errors.New("empty password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func comparePW(password, hashedPW string) bool {
	return password == hashedPW
}
