package article

import (
	"articles-test-server/shared/api/usecase"
	"database/sql"
	"fmt"
)

type Usecase struct {
	usecase.BaseUsecase
	db         *sql.DB
	repository RepositoryInterface
}

// NewUsecase responses new Usecase instance.
func NewUsecase(bu *usecase.BaseUsecase, master *sql.DB, r RepositoryInterface) *Usecase {
	return &Usecase{BaseUsecase: *bu, db: master, repository: r}
}

type UsecaseInterface interface {
	GetArticles(page int) ([]Article, error)
}

func (u *Usecase) GetArticles(page int) ([]Article, error) {
	u.Logger.Info("start usecase operations")
	articles, err := u.repository.GetArticles(page)
	if err != nil {
		fmt.Println("error rep functions", err)
	}
	return articles, nil
}
