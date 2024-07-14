package middleware

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	logger "github.com/HwaI12/go-api-tutorial/internal/log"
	"github.com/HwaI12/go-api-tutorial/internal/transaction"
)

// モックのハンドラーを作成する
func init() {
	logger.InitializeLogger()
	transaction.InitializeGlobalTransaction()
}

func TestAPIKeyAuthMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		apiKey         string
		expectedStatus int
	}{
		{
			name:           "API key is empty",
			apiKey:         "",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "API key is invalid",
			apiKey:         "invalid_api_key",
			expectedStatus: http.StatusUnauthorized,
		},
		{
			name:           "API key is valid",
			apiKey:         "valid_api_key",
			expectedStatus: http.StatusOK,
		},
	}

	// APIキーを環境変数に設定
	expectedAPIKey := "valid_api_key"
	os.Setenv("API_KEY", expectedAPIKey)

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// トランザクションの初期化
			ctx := context.Background()
			ctx = transaction.InitializeTransaction(ctx)

			// リクエストを作成
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			req = req.WithContext(ctx)
			if tt.apiKey != "" {
				req.Header.Set("X-API-KEY", tt.apiKey)
			}

			// レスポンスを作成
			rr := httptest.NewRecorder()

			// 次のハンドラーを作成
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			// テスト対象のミドルウェアを作成
			handler := APIKeyAuthMiddleware(next)
			handler.ServeHTTP(rr, req)

			// ステータスコードが期待通りか確認
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}

func TestTransactionMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "Transaction context is set",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// トランザクションの初期化
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rr := httptest.NewRecorder()

			// モックのハンドラーを作成
			// トランザクションIDとトランザクション時間がコンテキストに設定されているかどうかを確認する
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				trnID := r.Context().Value(transaction.TrnIDKey)
				trnTime := r.Context().Value(transaction.TrnTimeKey)

				if trnID == nil || trnTime == nil {
					http.Error(w, "Transaction context not set", http.StatusInternalServerError)
					return
				}

				w.WriteHeader(http.StatusOK)
			})

			// テスト対象のミドルウェアを作成
			handler := TransactionMiddleware(next)
			handler.ServeHTTP(rr, req)

			// ステータスコードが期待通りか確認
			// トランザクション情報が正しく設定されている場合は http.StatusOK を返し、設定されていない場合はエラーメッセージを返す
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}
		})
	}
}
