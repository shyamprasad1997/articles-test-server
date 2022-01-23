package middlewares

import (
	"articles-test-server/shared/api/repository"
	"database/sql"

	"github.com/sirupsen/logrus"
)

// Repository struct.
type AuthenticationRepository struct {
	repository.BaseRepository
	// connect master database.
	masterDB *sql.DB
	// connect read replica database.
	readDB *sql.DB
}

type IAuthenticationRepository interface {
	CheckIfValidUser(authorID int, authorEmail string) error
	CheckIfUserIsAdmin(userId int) error
}

//CheckIfValidUser- Checks with the db whether the user is valid
func (r *AuthenticationRepository) CheckIfValidUser(authorID int, authorEmail string) error {
	tx, err := r.readDB.Begin()
	if err != nil {
		logrus.Warn("Error in respository.CheckIfValidUser(): ", err)
		return err
	}
	query := `SELECT email, user_id from user_app where email=$1 AND user_id=$2`
	err = tx.QueryRow(query, authorEmail, authorID).Scan(&authorEmail, &authorID)
	tx.Commit()
	return err
}

//CheckIfUserIsAdmin- Checks with the db whether the user is admin
func (r *AuthenticationRepository) CheckIfUserIsAdmin(userId int) error {
	tx, err := r.readDB.Begin()
	if err != nil {
		logrus.Warn("Error in respository.CheckIfUserIsAdmin(): ", err)
		return err
	}
	query := `SELECT user_id from user_app where user_id=$1 AND is_admin=true`
	err = tx.QueryRow(query, userId).Scan(&userId)
	tx.Commit()
	return err
}

func NewAuthenticationRepository(master, read *sql.DB) IAuthenticationRepository {
	return &AuthenticationRepository{masterDB: master, readDB: read}
}
