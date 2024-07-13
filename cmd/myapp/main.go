package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/HwaI12/go-api-tutorial/api"
	"github.com/HwaI12/go-api-tutorial/internal/logger"
	"github.com/HwaI12/go-api-tutorial/internal/middleware"
	"github.com/HwaI12/go-api-tutorial/internal/transaction"
	"github.com/HwaI12/go-api-tutorial/pkg/database"
)

func main() {
	// ロガーの初期化
	logger.InitializeLogger()

	// グローバルトランザクションの初期化
	transaction.InitializeGlobalTransaction()

	// トランザクションの初期化
	ctx := context.Background()
	ctx = transaction.InitializeTransaction(ctx)

	// ログエントリの作成とトランザクション情報の追加
	entry := logger.WithTransaction(ctx)

	entry.Info(".envファイルの読み込みを開始します")
	err := godotenv.Load()
	if err != nil {
		entry.Error(".envファイルの読み込みに失敗しました")
	} else {
		entry.Info(".envファイルの読み込みに成功しました")
	}

	entry.Info("データベースに接続します")
	db, err := database.Connect(ctx)
	if err != nil {
		entry.WithError(err).Fatal("データベースへの接続に失敗しました")
	} else {
		entry.Info("データベースに接続しました")
	}

	entry.Info("ルーティングを設定します")
	router := mux.NewRouter()
	router.Use(middleware.TransactionMiddleware) // トランザクションミドルウェアを使用
	router.Use(middleware.APIKeyAuthMiddleware)  // APIキー認証ミドルウェアを使用
	api.RegisterRoutes(router, db)

	// サーバーシャットダウンの処理
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// サーバー終了時にログを出力
	defer entry.Info("サーバーが終了しました")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			entry.WithError(err).Fatal("サーバーの起動に失敗しました")
		}
	}()
	entry.Info("サーバーが正常に起動しました")
	fmt.Printf("http://localhost:8080 でサーバーが起動しました\n")

	// シグナルを待つ
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// サーバーのシャットダウン
	entry.Info("サーバーのシャットダウンを開始します")
	if err := server.Shutdown(context.Background()); err != nil {
		entry.WithError(err).Fatal("サーバーのシャットダウンに失敗しました")
	}
	entry.Info("サーバーのシャットダウンが完了しました")
}
