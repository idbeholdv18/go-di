package middleware

import (
	"github/idbeholdv18/go-di/internal/domain"
	"net/http"
	"strings"
)

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			domain.SendError(w, r, domain.NewUnauthorizedError)
			return
		}

		token := strings.TrimPrefix(authHeader, "Basic")
		if token == "" {
			domain.SendError(w, r, domain.NewUnauthorizedError)
			return
		}

		h(w, r)
	})
}
