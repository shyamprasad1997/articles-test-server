package middlewares

import (
	"database/sql"
	"errors"

	config "articles-test-server/configs"
	"articles-test-server/shared/api/repository"
	"articles-test-server/shared/utils"

	"github.com/dgrijalva/jwt-go"
)

type AuthenticationService struct {
	IAuthenticationRepository
}

type IAuthenticationService interface {
	CheckIfTokenValid(reqToken string) error
	CheckIfTokenAdmin(reqToken string) error
}

// CheckIfTokenValid- checks whether the token is a valid
func (service *AuthenticationService) CheckIfTokenValid(reqToken string) error {
	config := config.Load()
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Login.Token), nil
	})

	if err != nil || !token.Valid || token == nil {
		return errors.New("invalid token")
	} else {
		authorID := int(claims["id"].(float64))
		authorEmail := claims["email"].(string)
		err := service.CheckIfValidUser(authorID, authorEmail)
		if err != nil {
			return errors.New("unauthorized user")
		}
	}
	return nil
}

// CheckIfTokenAdmin- checks whether the token is an admin
func (service *AuthenticationService) CheckIfTokenAdmin(reqToken string) error {
	userId, err := utils.GetUserId(reqToken)
	if err != nil || userId <= 0 {
		return err
	}
	err = service.CheckIfUserIsAdmin(userId)
	return err
}

func NewAuthenticationService(br *repository.BaseRepository, master, read *sql.DB) IAuthenticationService {
	return &AuthenticationService{NewAuthenticationRepository(master, read)}
}
