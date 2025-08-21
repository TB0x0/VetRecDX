package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func checkEmpty(input, field string) error {
	if input == "" {
		return fmt.Errorf("missing %s", field)
	}
	return nil
}

func generateClaims(authConfig AuthConfig, userEntry UserDBEntry) (jwt.MapClaims, error) {
	//collapse all errors into single output
	var errorsList []string

	if authConfig.SecretKey == nil {
		errorsList = append(errorsList, "missing secret key")
	}
	if err := checkEmpty(userEntry.Id, "user ID"); err != nil {
		errorsList = append(errorsList, err.Error())
	}
	if err := checkEmpty(userEntry.Username, "username"); err != nil {
		errorsList = append(errorsList, err.Error())
	}
	if err := checkEmpty(authConfig.Issuer, "issuer"); err != nil {
		errorsList = append(errorsList, err.Error())
	}

	if len(errorsList) > 0 {
		return nil, errors.New(strings.Join(errorsList, "; "))
	}

	claims := jwt.MapClaims{
		"sub":          userEntry.Id,
		"username":     userEntry.Username,
		"passwordhash": userEntry.PasswordHash,
		"role":         userEntry.Role,
		"iss":          authConfig.Issuer,
		"iat":          time.Now().Unix(),
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	}

	return claims, nil
}

func encodeJWT(authConfig AuthConfig, userEntry UserDBEntry) (string, error) {
	claims, err := generateClaims(authConfig, userEntry)
	if err != nil {
		return "", errors.New("error generating claim")
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	newTokenString, err := newToken.SignedString(authConfig.SecretKey)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return newTokenString, nil
}
