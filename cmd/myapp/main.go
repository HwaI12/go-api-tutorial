package main

import (
	"context"
	"go-api-tutorial/api"
	"go-api-tutorial/internal/Log"
	"go-api-tutorial/internal/transaction"
	"go-api-tutorial/pkg/database"

	"github.com/sirupsen/logrus"

	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

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

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := database.Connect()
	router := mux.NewRouter()
	api.RegisterRoutes(router, db)
	log.Fatal(http.ListenAndServe(":8080", router))
}
