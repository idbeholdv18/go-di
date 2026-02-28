package middleware

import (
	"github/idbeholdv18/go-di/internal/domain"
	"github/idbeholdv18/go-di/internal/service"
	"net/http"
	"strings"
)

func AuthMiddleware(authService service.AuthService) func(h http.HandlerFunc) http.HandlerFunc {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				domain.SendError(w, r, domain.NewUnauthorizedError)
				return
			}

			token := strings.TrimPrefix(authHeader, "Basic")
			if !authService.Validate(token) {
				domain.SendError(w, r, domain.NewUnauthorizedError)
				return
			}
			h(w, r)
		}
	}
}
