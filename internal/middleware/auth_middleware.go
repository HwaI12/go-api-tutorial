package middleware

import (
	"net/http"
	"os"

	"github.com/HwaI12/go-api-tutorial/internal/errors"
	"github.com/HwaI12/go-api-tutorial/internal/logger"
	"github.com/HwaI12/go-api-tutorial/internal/transaction"
	"github.com/HwaI12/go-api-tutorial/internal/views"
)

// APIKeyAuthMiddlewareはAPIキー認証を行うミドルウェア
func APIKeyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // リクエストのコンテキストを使用
		entry := logger.WithTransaction(ctx)

		apiKey := r.Header.Get("X-API-KEY")

		// APIキーが空の場合はエラーレスポンスを返す
		if apiKey == "" {
			entry.Error("APIキーが空です")
			views.RespondWithError(w, ctx, errors.APIKeyEmptyError())
			return
		}

		// APIキーが期待される値と異なる場合はエラーレスポンスを返す
		expectedAPIKey := os.Getenv("API_KEY")
		if apiKey != expectedAPIKey {
			entry.Error("APIキーが無効です")
			views.RespondWithError(w, ctx, errors.InvalidAPIKeyError())
			return
		}

		next.ServeHTTP(w, r)
	})
}

// トランザクション情報をコンテキストに設定するミドルウェア
func TransactionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := transaction.InitializeTransaction(r.Context())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
