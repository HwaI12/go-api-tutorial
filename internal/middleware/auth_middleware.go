package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/HwaI12/go-api-tutorial/internal/errors"
	"github.com/HwaI12/go-api-tutorial/internal/views"
)

// APIKeyAuthMiddlewareはAPIキー認証を行うミドルウェアである
func APIKeyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background() // 新しいコンテキストを作成
		apiKey := r.Header.Get("X-API-KEY")

		// APIキーが空の場合はエラーレスポンスを返す
		if apiKey == "" {
			views.RespondWithError(w, ctx, errors.InvalidAPIKeyError())
			return
		}

		// APIキーが期待される値と異なる場合はエラーレスポンスを返す
		expectedAPIKey := os.Getenv("API_KEY")
		if apiKey != expectedAPIKey {
			views.RespondWithError(w, ctx, errors.InvalidAPIKeyError())
			return
		}

		next.ServeHTTP(w, r)
	})
}
