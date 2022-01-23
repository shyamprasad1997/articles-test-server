package article

import (
	"articles-test-server/shared/api/repository"
	"database/sql"
)

// Repository struct.
type Repository struct {
	repository.BaseRepository
	// connect master database.
	masterDB *sql.DB
	// connect read replica database.
	readDB *sql.DB
}

// NewRepository responses new Repository instance.
func NewRepository(br *repository.BaseRepository, master *sql.DB, read *sql.DB) *Repository {
	return &Repository{BaseRepository: *br, masterDB: master, readDB: read}
}

type RepositoryInterface interface {
	GetArticles(page int) ([]Article, error)
}

func (r *Repository) GetArticles(page int) ([]Article, error) {
	articles := []Article{}
	r.Logger.Info("Fetch article details")
	query := `SELECT id_article, title, "desc" from articles LIMIT 10 OFFSET $1`
	rows, err := r.readDB.Query(query, page)
	if err != nil {
		return articles, err
	}
	for rows.Next() {
		var article Article
		rows.Scan(&article.ArticleId, &article.Title, &article.Description)
		articles = append(articles, article)
	}
	return articles, nil
}
