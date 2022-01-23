package handler

import (
	"articles-test-server/infrastructure"
	"articles-test-server/shared/utils"
	"net/http"
)

// BaseHTTPHandler base handler struct.
type BaseHTTPHandler struct {
	Logger infrastructure.LoggerInterface
}

// ResponseJSON responses status code 200 and json.
func (h *BaseHTTPHandler) ResponseJSON(w http.ResponseWriter, data interface{}) {
	// status code 200
	utils.ResponseJSON(w, http.StatusOK, data)
}

// StatusBadRequest responses status code 400 and json.
func (h *BaseHTTPHandler) StatusBadRequest(w http.ResponseWriter, data interface{}) {
	// status code 400
	utils.ResponseJSON(w, http.StatusBadRequest, data)
}

// StatusUnauthorized responses status code 401 and json.
func (h *BaseHTTPHandler) StatusUnauthorized(w http.ResponseWriter, data interface{}) {
	// status code 401
	utils.ResponseJSON(w, http.StatusUnauthorized, data)
}

// StatusServerError responses 500.
func (h *BaseHTTPHandler) StatusServerError(w http.ResponseWriter, data interface{}) {
	// status code 500
	utils.ResponseJSON(w, http.StatusInternalServerError, data)
}

// NewBaseHTTPHandler returns BaseHTTPHandler instance.
func NewBaseHTTPHandler(logger infrastructure.LoggerInterface) *BaseHTTPHandler {
	return &BaseHTTPHandler{Logger: logger}
}
