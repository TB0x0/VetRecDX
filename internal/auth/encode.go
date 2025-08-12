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

func generateClaims(userInfo User, authConfig AuthConfig) (jwt.MapClaims, error) {
	//collapse all errors into single output
	var errorsList []string

	if err := checkEmpty(userInfo.Id, "user ID"); err != nil {
		errorsList = append(errorsList, err.Error())
	}
	if err := checkEmpty(userInfo.Email, "email"); err != nil {
		errorsList = append(errorsList, err.Error())
	}
	if err := checkEmpty(authConfig.Issuer, "issuer"); err != nil {
		errorsList = append(errorsList, err.Error())
	}
	if authConfig.SecretKey == nil {
		errorsList = append(errorsList, "missing secret key")
	}

	if len(errorsList) > 0 {
		return nil, errors.New(strings.Join(errorsList, "; "))
	}

	claims := jwt.MapClaims{
		"sub":         userInfo.Id,
		"email":       userInfo.Email,
		"roles":       userInfo.Roles,
		"permissions": userInfo.Permissions,
		"org_id":      userInfo.Org_id,
		"iss":         authConfig.Issuer,
		"iat":         time.Now().Unix(),
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	return claims, nil
}

func encodeJWT(userInfo User, authConfig AuthConfig) (string, error) {
	claims, err := generateClaims(userInfo, authConfig)
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
