package middleware

import (
	"go-api-tutorial/internal/views"
	"net/http"
	"os"

	"github.com/HwaI12/go-api-tutorial/internal/errors"
)

// APIKeyAuthMiddlewareはAPIキー認証を行うミドルウェアである
func APIKeyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" {
			views.RespondWithError(w, errors.InvalidAPIKeyError())
			return
		}

		expectedAPIKey := os.Getenv("API_KEY")
		if apiKey != expectedAPIKey {
			views.RespondWithError(w, errors.InvalidAPIKeyError())
			return
		}

		next.ServeHTTP(w, r)
	})
}
