package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	errors "github.com/HwaI12/go-api-tutorial/internal/error"
	logger "github.com/HwaI12/go-api-tutorial/internal/log"
)

type Book struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type BookInput struct {
	Name  *string `json:"name"`
	Price *int    `json:"price"`
}

// Validate は Book モデルの検証を行う
func (b *Book) Validate(ctx context.Context) error {
	entry := logger.WithTransaction(ctx)

	if b.Name == "" {
		entry.Errorf("パラメータ'name'が空です。本の名前を入力してください")
		return errors.BookNameEmptyError()
	}

	if b.Price == 0 {
		entry.Errorf("パラメータ'price'が0です。本の価格を入力してください")
		return errors.BookPriceEmptyError()
	}

	if len(b.Name) > 50 {
		entry.Errorf("パラメータ'name'が長すぎます。50文字以内で書いてください")
		return errors.BookNameTooLongError()
	}
	entry.Infof("本の名前が50文字以内であることを確認しました")

	if b.Price < 0 {
		entry.Errorf("パラメータ'price'が0以下です。正の整数を入力してください")
		return errors.BookPriceNegativeError()
	}
	entry.Infof("本の価格が0以上であることを確認しました")

	if b.Price > 20000 {
		entry.Errorf("本の価格が高すぎます")
		return errors.BookPriceTooHighError()
	}
	entry.Infof("本の価格が20000円以下であることを確認しました")

	return nil
}

// GetBooks はデータベースから書籍を取得する
func GetBooks(ctx context.Context, db *sql.DB) ([]Book, error) {
	entry := logger.WithTransaction(ctx)

	entry.Infof("GetBooks関数が呼び出されました")

	rows, err := db.Query("SELECT id, name, price, created_at FROM books")
	if err != nil {
		entry.Errorf("データベースからの取得に失敗しました: %v", err)
		return nil, errors.DatabaseQueryError()
	}
	entry.Infof("データベースからの取得に成功しました")
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		err := rows.Scan(&book.ID, &book.Name, &book.Price, &book.CreatedAt)
		if err != nil {
			entry.Errorf("データベース結果のスキャンに失敗しました: %v", err)
			return nil, errors.DatabaseScanError()
		}
		books = append(books, book)
	}
	entry.Infof("データベース結果のスキャンに成功しました")

	entry.Infof("GetBooks関数が終了しました")
	return books, nil
}

// CreateBook はデータベースに書籍を登録する
func (b *Book) CreateBook(ctx context.Context, db *sql.DB) error {
	entry := logger.WithTransaction(ctx)

	entry.Infof("CreateBook関数が呼び出されました")

	stmt, err := db.Prepare("INSERT INTO books(name, price) VALUES(?, ?)")
	if err != nil {
		entry.Errorf("SQLステートメントの準備に失敗しました: %v", err)
		return errors.SQLPreparationError()
	}
	entry.Infof("SQLステートメントの準備に成功しました")
	defer stmt.Close()

	result, err := stmt.Exec(b.Name, b.Price)
	if err != nil {
		entry.Errorf("データベースへの挿入に失敗しました: %v", err)
		return errors.DatabaseInsertError()
	}
	entry.Infof("データベースへの挿入に成功しました")

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		entry.Errorf("最後に挿入されたIDの取得に失敗しました: %v", err)
		return errors.LastInsertIDError()
	}
	entry.Infof("最後に挿入されたIDの取得に成功しました")

	b.ID = fmt.Sprintf("%d", lastInsertId)
	b.CreatedAt = time.Now()

	entry.Infof("本の登録に成功しました")
	entry.Infof("CreateBook関数が終了しました")
	return nil
}
