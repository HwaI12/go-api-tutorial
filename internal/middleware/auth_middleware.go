package middleware

import (
	"context"
	"net/http"
	"os"

	"github.com/HwaI12/go-api-tutorial/internal/errors"
	"github.com/HwaI12/go-api-tutorial/internal/logger"
	"github.com/HwaI12/go-api-tutorial/internal/transaction"
	"github.com/HwaI12/go-api-tutorial/internal/views"
	"github.com/sirupsen/logrus"
)

// APIKeyAuthMiddlewareはAPIキー認証を行うミドルウェア
func APIKeyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // リクエストのコンテキストを使用
		entry := logger.WithTransaction(ctx)

		apiKey := r.Header.Get("X-API-KEY")

		// APIキーが空の場合はエラーレスポンスを返す
		if apiKey == "" {
			err := errors.APIKeyEmptyError()
			entry.WithError(err).Error("APIキーが空です")
			logAndRespondWithError(w, ctx, entry, err)
			return
		}

		// APIキーが期待される値と異なる場合はエラーレスポンスを返す
		expectedAPIKey := os.Getenv("API_KEY")
		if apiKey != expectedAPIKey {
			err := errors.InvalidAPIKeyError()
			entry.WithError(err).Error("APIキーが無効です")
			logAndRespondWithError(w, ctx, entry, err)
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

// logAndRespondWithError はエラーレスポンスを返し、エラーメッセージをログに記録する
func logAndRespondWithError(w http.ResponseWriter, ctx context.Context, entry *logrus.Entry, err *errors.UserDefinedError) {
	response := views.CreateExceptionResponse(ctx, err)
	entry.Debugf("%+v", err)
	entry.Debugf("レスポンス結果: %+v", response)
	views.RespondWithError(w, ctx, err)
}
