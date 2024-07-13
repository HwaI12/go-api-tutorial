package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/HwaI12/go-api-tutorial/api"
	"github.com/HwaI12/go-api-tutorial/internal/Log"
	"github.com/HwaI12/go-api-tutorial/internal/transaction"
	"github.com/HwaI12/go-api-tutorial/pkg/database"
)

func main() {
	// ロガーの初期化
	Log.InitializeLogger()

	// トランザクションの初期化
	ctx := context.Background()
	ctx = transaction.InitializeTransaction(ctx)

	// ログエントリの作成とトランザクション情報の追加
	entry := logrus.WithContext(ctx)
	entry = entry.WithFields(logrus.Fields{
		"trn_id":   ctx.Value(transaction.TrnIDKey),
		"trn_time": ctx.Value(transaction.TrnTimeKey),
	})

	entry.Info(".envファイルの読み込みを開始します")
	err := godotenv.Load()
	if err != nil {
		entry.Error(".envファイルの読み込みに失敗しました")
	} else {
		entry.Info(".envファイルの読み込みに成功しました")
	}

	entry.Info("データベースに接続します")
	db, err := database.Connect()
	if err != nil {
		entry.WithError(err).Fatal("データベースへの接続に失敗しました")
	} else {
		entry.Info("データベースに接続しました")
	}

	entry.Info("ルーティングを設定します")
	router := mux.NewRouter()
	api.RegisterRoutes(router, db)

	entry.Info("サーバーを起動します")
	if err := http.ListenAndServe(":8080", router); err != nil {
		entry.WithError(err).Fatal("サーバーの起動に失敗しました")
	}
}
