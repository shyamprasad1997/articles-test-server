package article

import (
	"articles-test-server/shared/api/handler"
	"articles-test-server/shared/api/repository"
	"articles-test-server/shared/api/usecase"
	"articles-test-server/shared/utils"

	"database/sql"
	"encoding/json"
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
		h.Logger.Warn(err, "failed to fetch")
		response.Status = false
		h.StatusServerError(w, response)
		return
	}
	response.Status = true
	response.Articles = articles
	h.Logger.Info("successfully returned")
	h.ResponseJSON(w, response)
}

func (h *HTTPHandler) GetSearchArticlesHandler(w http.ResponseWriter, r *http.Request) {
	var response GetArticlesResponse
	pageParam := chi.URLParam(r, "page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		h.Logger.Warn(err, "bad request")
		response.Status = false
		h.StatusBadRequest(w, response)
		return
	}
	queryValues := r.URL.Query()
	key := queryValues.Get("key")
	if key == "" {
		h.Logger.Warn(err, "bad request")
		response.Status = false
		h.StatusBadRequest(w, response)
		return
	}
	articles, err := h.usecase.SearchArticles(page-1, key)
	if err != nil {
		h.Logger.Warn(err, "search failed")
		response.Status = false
		h.StatusServerError(w, response)
		return
	}
	response.Status = true
	response.Articles = articles
	h.Logger.Info("successfully returned")
	h.ResponseJSON(w, response)
}

func (h *HTTPHandler) PostArticle(w http.ResponseWriter, r *http.Request) {
	var request PostArticlesRequest
	var response PostArticlesResponse
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		h.Logger.Warn(err, "bad request")
		h.StatusBadRequest(w, response)
		return
	}
	reqToken := r.Header.Get("Authorization")
	userId, err := utils.GetUserId(reqToken)
	if err != nil || userId <= 0 {
		h.Logger.Warn(err, "bad request")
		h.StatusBadRequest(w, response)
		return
	}
	request.CreatedBy = uint64(userId)
	err = h.usecase.CreateArticle(request)
	if err != nil {
		h.Logger.Warn(err, "failed to create")
		h.StatusServerError(w, response)
		return
	}
	response.Status = true
	h.Logger.Info("successfully created")
	h.ResponseJSON(w, response)
}

func (h *HTTPHandler) ApproveArticle(w http.ResponseWriter, r *http.Request) {
	var response PostArticlesResponse
	articleIdParam := chi.URLParam(r, "article_id")
	articleID, err := strconv.Atoi(articleIdParam)
	if err != nil || articleID < 1 {
		h.Logger.Warn(err, "bad request")
		response.Status = false
		h.StatusBadRequest(w, response)
		return
	}
	err = h.usecase.ApproveArticle(articleID)
	if err != nil {
		h.Logger.Warn(err, "failed to approve")
		h.StatusServerError(w, response)
		return
	}
	response.Status = true
	h.Logger.Info("successfully approved")
	h.ResponseJSON(w, response)
}

func (h *HTTPHandler) DeclineArticle(w http.ResponseWriter, r *http.Request) {
	var response PostArticlesResponse
	articleIdParam := chi.URLParam(r, "article_id")
	articleID, err := strconv.Atoi(articleIdParam)
	if err != nil || articleID < 1 {
		h.Logger.Warn(err, "bad request")
		response.Status = false
		h.StatusBadRequest(w, response)
		return
	}
	err = h.usecase.DeclineArticle(articleID)
	if err != nil {
		h.Logger.Warn(err, "failed to approve")
		h.StatusServerError(w, response)
		return
	}
	response.Status = true
	h.Logger.Info("successfully declined")
	h.ResponseJSON(w, response)
}
