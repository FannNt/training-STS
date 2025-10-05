package middleware

import (
	"net/http"

	"book-api/storage"
	"book-api/utils"
)

// AuthMiddleware creates an authentication middleware
func AuthMiddleware(sessionStorage *storage.SessionStorage) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from Authorization header
			token := r.Header.Get("Authorization")
			if token == "" {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, "Authorization token is required")
				return
			}

			// Remove "Bearer " prefix if present
			if len(token) > 7 && token[:7] == "Bearer " {
				token = token[7:]
			}

			// Validate token
			_, valid := sessionStorage.ValidateToken(token)
			if !valid {
				utils.WriteErrorResponse(w, http.StatusUnauthorized, "Invalid or expired token")
				return
			}

			// Token is valid, proceed to next handler
			next.ServeHTTP(w, r)
		})
	}
}
