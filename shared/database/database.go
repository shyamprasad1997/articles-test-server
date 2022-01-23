package database

import (
	"articles-test-server/shared/utils"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

// ErrNoRecord error
var ErrNoRecord = errors.New("No Records From Database")

//ConnectDB connects master db and read db to the application
func ConnectDB() (*sql.DB, *sql.DB, error) {
	masterDB, err := ConnectMasterDB()
	if err != nil {
		return nil, nil, err
	}
	readDB, err := ConnectReadDB()
	if err != nil {
		return nil, nil, err
	}
	return masterDB, readDB, nil
}

//ConnectMasterDB connects master db to application
func ConnectMasterDB() (*sql.DB, error) {
	masterStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		utils.Env.DbMasterHost, utils.Env.DbPort, utils.Env.DbMasterUser, utils.Env.DbMasterPassword, utils.Env.DbName)

	masterDB, err := sql.Open("postgres", masterStr)
	if err != nil {
		return nil, utils.ErrorsWrap(err, "error in connecting master db")
	}
	return masterDB, nil
}

//ConnectReadDB connects read db to application
func ConnectReadDB() (*sql.DB, error) {
	readStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		utils.Env.DbReadHost, utils.Env.DbPort, utils.Env.DbReadUser, utils.Env.DbReadPassword, utils.Env.DbName)

	readDB, err := sql.Open("postgres", readStr)
	if err != nil {
		return nil, utils.ErrorsWrap(err, "error in connecting read db")
	}
	return readDB, nil
}
