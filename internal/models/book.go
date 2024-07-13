package models

import (
	"database/sql"
	"strconv"
	"time"
)

type Book struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllBooks(db *sql.DB) ([]Book, error) {
	// データベースから全ての本を取得するSQLクエリ
	query := "SELECT * FROM books"

	// データベースにクエリを送信し、結果を取得
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	// 関数が終了した際にrowsを閉じる
	defer rows.Close()

	// 取得した本を格納するスライス
	books := []Book{}

	// 取得したデータをスライスに格納
	for rows.Next() {
		var book Book
		// 取得したデータをbook構造体に格納
		if err := rows.Scan(&book.ID, &book.Name, &book.Price, &book.CreatedAt); err != nil {
			return nil, err
		}
		// スライスに本を追加
		books = append(books, book)
	}

	// エラーが発生した場合はエラーを返す
	return books, nil
}

func (b *Book) CreateBook(db *sql.DB) error {
	// データベースに新しい本を追加するためのSQLクエリ
	stmt, err := db.Prepare("INSERT INTO books(name, price) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer db.Close()

	// データベースに新しい本を追加
	result, err := stmt.Exec(b.Name, b.Price)
	if err != nil {
		return err
	}

	// 最後に追加された本のIDを取得
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// 本のIDを構造体に設定
	// strconv.FormatIntはint64型を文字列に変換する関数
	// なぜなら、
	b.ID = strconv.FormatInt(lastInsertId, 10)

	// 本の作成日時を設定
	b.CreatedAt = time.Now()

	// エラーが発生した場合はエラーを返す
	return nil
}
