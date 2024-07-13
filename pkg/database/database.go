package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf(".envファイルの読み込みに失敗しました: %v", err)
	}

	// 環境変数が設定されているか確認
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return nil, fmt.Errorf("必要な環境変数が設定されていません")
	}

	// データベース接続文字列を作成
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Printf("データベース接続文字列: %s\n", dataSourceName) // デバッグ用

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("sql.Openによるデータベース接続に失敗しました: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db.PingによるデータベースへのPingに失敗しました: %v", err)
	}

	return db, nil
}
