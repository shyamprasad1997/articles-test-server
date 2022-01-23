package article

import (
	"articles-test-server/shared/api/handler"
	"articles-test-server/shared/api/repository"
	"articles-test-server/shared/api/usecase"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.BaseHTTPHandler
	usecase UsecaseInterface
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository, read, master *sql.DB) *HTTPHandler {
	or := NewRepository(br, master, read)
	ou := NewUsecase(bu, master, or)
	return &HTTPHandler{BaseHTTPHandler: *bh, usecase: ou}
}

func (h *HTTPHandler) GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	var response GetArticlesResponse
	pageParam := chi.URLParam(r, "page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		h.Logger.Warn(err, "bad request")
		response.Status = false
		h.StatusBadRequest(w, response)
		return
	}
	articles, err := h.usecase.GetArticles(page - 1)
	if err != nil {
		h.Logger.Warn(err, "bad request")
		response.Status = false
		h.StatusServerError(w, response)
		return
	}
	response.Status = true
	response.Articles = articles
	h.Logger.Info("successfully returned")
	h.ResponseJSON(w, response)
}
