package article

import (
	"articles-test-server/shared/api/usecase"
	"database/sql"
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
	SearchArticles(page int, key string) ([]Article, error)
	CreateArticle(request PostArticlesRequest) error
	ApproveArticle(articleId int) error
	DeclineArticle(articleId int) error
}

func (u *Usecase) GetArticles(page int) ([]Article, error) {
	u.Logger.Info("start usecase operations - GetArticles()")
	articles, err := u.repository.GetArticles(page)
	if err != nil {
		u.Logger.Warn("error repo functions", err)
	}
	return articles, nil
}

func (u *Usecase) SearchArticles(page int, key string) ([]Article, error) {
	u.Logger.Info("start usecase operations - SearchArticles()")
	articles, err := u.repository.SearchArticles(page, key)
	if err != nil {
		u.Logger.Warn("error repo functions", err)
	}
	return articles, nil
}

func (u *Usecase) CreateArticle(request PostArticlesRequest) error {
	u.Logger.Info("start usecase operations - CreateArticle()")
	err := u.repository.CreateArticle(request)
	if err != nil {
		u.Logger.Warn("error repo functions", err)
	}
	return nil
}

func (u *Usecase) ApproveArticle(articleId int) error {
	u.Logger.Info("start usecase operations - ApproveArticle()")
	err := u.repository.ApproveArticle(articleId)
	if err != nil {
		u.Logger.Warn("error repo functions", err)
	}
	return nil
}

func (u *Usecase) DeclineArticle(articleId int) error {
	u.Logger.Info("start usecase operations - DeclineArticle()")
	err := u.repository.DeclineArticle(articleId)
	if err != nil {
		u.Logger.Warn("error repo functions", err)
	}
	return nil
}
