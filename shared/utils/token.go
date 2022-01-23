package utils

import (
	config "articles-test-server/configs"

	"github.com/dgrijalva/jwt-go"
)

// GetUserId - to find user ID
func GetUserId(token string) (int, error) {
	config := config.Load()
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Login.Token), nil
	})
	if err != nil {
		return 0, err
	}
	author := int(claims["id"].(float64))
	return author, nil
}
