package router

import (
	"articles-test-server/article"
	"articles-test-server/infrastructure"
	"articles-test-server/middlewares"
	"articles-test-server/shared/api/handler"
	"articles-test-server/shared/api/repository"
	"articles-test-server/shared/api/usecase"
	"database/sql"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router is application struct hold Mux and db connection
type Router struct {
	Mux           *chi.Mux
	MasterDB      *sql.DB
	ReadDB        *sql.DB
	LoggerHandler *infrastructure.Logger
}

// InitializeRouter initializes Mux and middleware
func (r *Router) InitializeRouter() {
	r.Mux.Use(middleware.RequestID)
	r.Mux.Use(middleware.RealIP)
	r.Mux.Use(middleware.Logger)
	r.Mux.Use(middleware.Recoverer)
	r.Mux.Use(middleware.Timeout(60 * time.Second))
}

// SetupHandler set database and redis and usecase.
func (r *Router) SetupHandler() {
	// base set.
	bh := handler.NewBaseHTTPHandler(r.LoggerHandler.Log)
	// base set.
	br := repository.NewBaseRepository(r.LoggerHandler.Log)
	// base set.
	bu := usecase.NewBaseUsecase(r.LoggerHandler.Log)

	ah := article.NewHTTPHandler(bh, bu, br, r.ReadDB, r.MasterDB)

	mw := middlewares.NewMiddleware(r.ReadDB, r.MasterDB, bh, bu, br)
	r.Mux.Route("/v1", func(cr chi.Router) {
		cr.Get("/articles/{page}", ah.GetArticlesHandler)
		cr.Get("/articles/search/{page}", ah.GetSearchArticlesHandler)
		cr.Group(func(subCr chi.Router) {
			subCr.Use(mw.TokenValidation)
			subCr.Post("/article", ah.PostArticle)
			subCr.Group(func(subCr chi.Router) {
				subCr.Use(mw.CheckIfAdmin)
				subCr.Put("/article/approve/{article_id}", ah.ApproveArticle)
				subCr.Put("/article/decline/{article_id}", ah.DeclineArticle)
			})
		})
	})
}
