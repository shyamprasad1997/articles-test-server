package main

import (
	"articles-test-server/shared/database"
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("wefwef")
	masterDB, _, err := database.ConnectDB()
	if err != nil {
		fmt.Println("Failed to initialize DB")
		return
	}
	fmt.Println("Starting Migration")
	err = StartMigration(masterDB)
	if err != nil {
		fmt.Println("Failed to Complete migration")
	}
}

func StartMigration(masterDB *sql.DB) error {
	tx, err := masterDB.Begin()
	if err != nil {
		fmt.Println("Failed to create transaction")
	}
	fmt.Println("Creating tables")
	err = CreateTables(tx)
	if err != nil {
		fmt.Println("Failed to create tables")
		tx.Rollback()
		return err
	}
	tx.Commit()
	fmt.Println("Created tables")
	tx, err = masterDB.Begin()
	if err != nil {
		fmt.Println("Failed to create transaction")
	}
	fmt.Println("Adding Data")
	err = AddUserData(tx)
	if err != nil {
		fmt.Println("Failed to add user data")
		return err
	}
	tx, err = masterDB.Begin()
	if err != nil {
		fmt.Println("Failed to create transaction")
	}
	err = AddArticleData(tx)
	if err != nil {
		fmt.Println("Failed to add article data")
		return err
	}
	fmt.Println("Added Data")
	return nil
}
