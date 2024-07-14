package models

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/HwaI12/go-api-tutorial/internal/errors"
	"github.com/HwaI12/go-api-tutorial/internal/logger"
)

type Book struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Book) Validate(ctx context.Context) error {
	entry := logger.WithTransaction(ctx)

	if b.Name == "" {
		entry.Errorf("本の名前が空です")
		return errors.BookNameEmptyError()
	}
	if len(b.Name) > 50 {
		entry.Errorf("本の名前が長すぎます")
		return errors.BookNameTooLongError()
	}
	if b.Price <= 0 {
		entry.Errorf("本の価格が空です")
		return errors.BookPriceEmptyError()
	}
	if b.Price > 20000 {
		entry.Errorf("本の価格が高すぎます")
		return errors.BookPriceTooHighError()
	}
	return nil
}

func (b *Book) CreateBook(ctx context.Context, db *sql.DB) error {
	entry := logger.WithTransaction(ctx)

	entry.Info("リクエストを受信しました")

	stmt, err := db.Prepare("INSERT INTO books(name, price) VALUES(?, ?)")
	if err != nil {
		entry.Errorf("SQLステートメントの準備に失敗しました: %v", err)
		return errors.SQLPreparationError()
	}
	defer stmt.Close()

	result, err := stmt.Exec(b.Name, b.Price)
	if err != nil {
		entry.Errorf("データベースへの挿入に失敗しました: %v", err)
		return errors.DatabaseInsertError()
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		entry.Errorf("最後に挿入されたIDの取得に失敗しました: %v", err)
		return errors.LastInsertIDError()
	}

	b.ID = strconv.FormatInt(lastInsertId, 10)
	b.CreatedAt = time.Now()

	entry.Infof("本の登録に成功しました: %+v", b)
	return nil
}
