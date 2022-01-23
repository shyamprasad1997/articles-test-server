package middlewares

import (
	"database/sql"
	"errors"

	config "articles-test-server/configs"
	"articles-test-server/shared/api/repository"

	"github.com/dgrijalva/jwt-go"
)

type AuthenticationService struct {
	IAuthenticationRepository
}

type IAuthenticationService interface {
	CheckIfTokenValid(reqToken string) error
}

// CheckIfTokenValid- checks whether the token is a valid
func (service *AuthenticationService) CheckIfTokenValid(reqToken string) error {
	config := config.Load()
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Login.Token), nil
	})

	if err != nil || !token.Valid || token == nil {
		return errors.New("Invalid token")
	} else {
		authorID := int(claims["id"].(float64))
		authorEmail := claims["email"].(string)
		err := service.CheckIfValidUser(authorID, authorEmail)
		if err != nil {
			return errors.New("Unauthorized user")
		}
	}
	return nil
}

func NewAuthenticationService(br *repository.BaseRepository, master, read *sql.DB) IAuthenticationService {
	return &AuthenticationService{NewAuthenticationRepository(master, read)}
}
