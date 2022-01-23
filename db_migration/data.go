package main

import (
	"database/sql"
	"fmt"
	"sync"
)

type User struct {
	Name    string
	Email   string
	IsAdmin bool
}

type Article struct {
	Title          string
	Description    string
	CreatedBy      int
	ApprovalStatus bool
	CreatedAt      string
}

func AddUserData(tx *sql.Tx) error {
	users := []User{
		{Name: "admin", Email: "admin@gmail.com", IsAdmin: true},
		{Name: "user", Email: "user@gmail.com", IsAdmin: true},
		{Name: "user1", Email: "user1@gmail.com", IsAdmin: true},
		{Name: "user2", Email: "user2@gmail.com", IsAdmin: true},
		{Name: "user3", Email: "user3@gmail.com", IsAdmin: true},
		{Name: "user4", Email: "user4@gmail.com", IsAdmin: true},
		{Name: "user5", Email: "user5@gmail.com", IsAdmin: true},
	}
	err := make(chan error, len(users))
	var wg sync.WaitGroup
	for _, user := range users {
		wg.Add(1)
		go AddUser(err, user, tx, &wg)
	}
	wg.Wait()
	if len(err) != 0 {
		fmt.Println("failed to complete data migration:", <-err)
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func AddUser(errCh chan error, user User, tx *sql.Tx, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := tx.Exec(`INSERT INTO user_app(name, email, is_admin) VALUES ($1, $2, $3);`, user.Name, user.Email, user.IsAdmin)
	if err != nil {
		errCh <- err
	}
}

func AddArticleData(tx *sql.Tx) error {
	articles := []Article{
		{Title: "title 1", Description: "This is my description.", ApprovalStatus: false, CreatedBy: 3, CreatedAt: "2022-01-23 17:46:18.524586+05:30"},
		{Title: "title 2", Description: "This is my description.", ApprovalStatus: true, CreatedBy: 3, CreatedAt: "2022-01-23 17:46:18.524586+05:30"},
		{Title: "title 3", Description: "This is my description.", ApprovalStatus: false, CreatedBy: 3, CreatedAt: "2022-01-23 17:46:18.524586+05:30"},
		{Title: "title 4", Description: "This is my description.", ApprovalStatus: false, CreatedBy: 3, CreatedAt: "2022-01-23 17:46:18.524586+05:30"},
		{Title: "title 5", Description: "This is my description.", ApprovalStatus: true, CreatedBy: 3, CreatedAt: "2022-01-23 17:46:18.524586+05:30"},
		{Title: "title 6", Description: "This is my description.", ApprovalStatus: true, CreatedBy: 3, CreatedAt: "2022-01-23 17:46:18.524586+05:30"},
		{Title: "title 7", Description: "This is my description.", ApprovalStatus: false, CreatedBy: 3, CreatedAt: "2022-01-23 17:46:18.524586+05:30"},
		{Title: "title 8", Description: "This is my description.", ApprovalStatus: true, CreatedBy: 3, CreatedAt: "2022-01-23 17:46:18.524586+05:30"},
		{Title: "title 9", Description: "This is my description.", ApprovalStatus: true, CreatedBy: 3, CreatedAt: "2022-01-23 17:46:18.524586+05:30"},
	}
	err := make(chan error, len(articles))
	var wg sync.WaitGroup
	for _, article := range articles {
		wg.Add(1)
		go AddArticle(err, article, tx, &wg)
	}
	wg.Wait()
	if len(err) != 0 {
		fmt.Println("failed to complete data migration:", <-err)
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func AddArticle(errCh chan error, article Article, tx *sql.Tx, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := tx.Exec(`INSERT INTO articles(title, "desc", created_at, created_by, approval_status) VALUES ($1, $2, $3, $4, $5);`, article.Title, article.Description, article.CreatedAt, article.CreatedBy, article.ApprovalStatus)
	if err != nil {
		errCh <- err
	}
}
