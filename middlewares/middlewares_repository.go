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
}

//CheckIfValidUser- Checks with the db whether the user is valid
func (r *AuthenticationRepository) CheckIfValidUser(authorID int, authorEmail string) error {
	tx, err := r.readDB.Begin()
	if err != nil {
		logrus.Warn("Error in respository.CheckIfValidUser(): ", err)
		return err
	}
	query := `SELECT email, user_id from user_app where email=$1 AND user_id=$2`
	err = tx.QueryRow(query, authorEmail, authorID).Scan()
	if err != sql.ErrNoRows {
		tx.Commit()
		return err
	}
	tx.Commit()
	return nil
}

func NewAuthenticationRepository(master, read *sql.DB) IAuthenticationRepository {
	return &AuthenticationRepository{masterDB: master, readDB: read}
}
