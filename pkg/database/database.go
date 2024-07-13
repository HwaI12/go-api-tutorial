package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/HwaI12/go-api-tutorial/internal/logger"
)

func Connect(ctx context.Context) (*sql.DB, error) {
	// トランザクション情報を含むロガーを取得
	entry := logger.WithTransaction(ctx)

	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		entry.WithError(err).Error(".envファイルの読み込みに失敗しました")
		return nil, fmt.Errorf(".envファイルの読み込みに失敗しました: %v", err)
	}

	// 環境変数が設定されているか確認
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		entry.Error("必要な環境変数が設定されていません")
		return nil, fmt.Errorf("必要な環境変数が設定されていません")
	}

	// データベース接続文字列を作成
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	entry.Info("データベース接続文字列: ", dataSourceName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		entry.WithError(err).Error("sql.Openによるデータベース接続に失敗しました")
		return nil, fmt.Errorf("sql.Openによるデータベース接続に失敗しました: %v", err)
	}

	if err := db.Ping(); err != nil {
		entry.WithError(err).Error("db.PingによるデータベースへのPingに失敗しました")
		return nil, fmt.Errorf("db.PingによるデータベースへのPingに失敗しました: %v", err)
	}

	entry.Info("データベース接続に成功しました")

	return db, nil
}
