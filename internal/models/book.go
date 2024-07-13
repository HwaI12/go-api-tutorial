package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Book struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (b *Book) Validate() error {
	if b.Name == "" {
		return errors.New("本の名前が空です")
	}
	if len(b.Name) > 50 {
		return errors.New("本の名前が長すぎます。50文字以内で書いてください")
	}
	if b.Price <= 0 {
		return errors.New("本の値段が空です。1文字以上書いてください")
	}
	if b.Price > 20000 {
		return errors.New("本の値段が高すぎます。20000円以内で書いてください")
	}
	return nil
}

func (b *Book) CreateBook(db *sql.DB) error {
	stmt, err := db.Prepare("INSERT INTO books(name, price) VALUES(?, ?)")
	if err != nil {
		return fmt.Errorf("SQLステートメントの準備に失敗しました: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(b.Name, b.Price)
	if err != nil {
		return fmt.Errorf("データベースへの挿入に失敗しました: %v", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("最後に挿入されたIDの取得に失敗しました: %v", err)
	}

	b.ID = strconv.FormatInt(lastInsertId, 10)
	b.CreatedAt = time.Now()

	return nil
}
