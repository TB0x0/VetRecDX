package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

func parseJWT(tokenString string, secretKey []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		},
	)

	if err != nil {
		return nil, errors.New("error parsing token")
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}

func extractUserInfo(token *jwt.Token) (*UserDBEntry, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token or unable to extract claims")
	}

	userInfo := &UserDBEntry{}

	if sub, ok := claims["sub"].(string); ok {
		userInfo.Id = sub
	}
	if username, ok := claims["username"].(string); ok {
		userInfo.Username = username
	}
	if passwordhash, ok := claims["passwordhash"].(string); ok {
		userInfo.PasswordHash = passwordhash
	}
	if role, ok := claims["role"].([]interface{}); ok {
		for _, r := range role {
			if str, ok := r.(string); ok {
				userInfo.Role = append(userInfo.Role, str)
			}
		}
	}

	return userInfo, nil
}

func decodeJWT(tokenString string, authConfig AuthConfig) (*UserDBEntry, error) {
	token, err := parseJWT(tokenString, authConfig.SecretKey)
	if err != nil {
		return nil, err
	}

	userInfo, err := extractUserInfo(token)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
