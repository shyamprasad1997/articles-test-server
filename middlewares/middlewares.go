package middlewares

import (
	"database/sql"
	"net/http"

	"articles-test-server/shared/api/handler"
	"articles-test-server/shared/api/repository"
	"articles-test-server/shared/api/usecase"

	logrus "github.com/sirupsen/logrus"
)

type AuthenticationMiddleware struct {
	handler.BaseHTTPHandler
	IAuthenticationService
}

// TokenValidation - Validate the given token
func (m *AuthenticationMiddleware) TokenValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		err := m.CheckIfTokenValid(reqToken)
		if err != nil {
			logrus.Warn(err)
			m.StatusUnauthorized(w, "Login Failed")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewMiddleware(read, master *sql.DB, bh *handler.BaseHTTPHandler, bu *usecase.BaseUsecase, br *repository.BaseRepository) *AuthenticationMiddleware {
	os := NewAuthenticationService(br, master, read)
	return &AuthenticationMiddleware{*bh, os}
}
