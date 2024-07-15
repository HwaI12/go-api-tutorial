package middleware

import (
	"net/http"
)

// CORS設定のミドルウェア
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-API-KEY")

		next.ServeHTTP(w, r)
	})
}
