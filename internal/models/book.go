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

// Validate は Book モデルの検証を行う
func (b *Book) Validate(ctx context.Context) error {
	entry := logger.WithTransaction(ctx)

	// パラメータの存在チェック
	params := map[string]interface{}{
		"name":  b.Name,
		"price": b.Price,
	}
	for param, value := range params {
		if value == nil || value == "" || value == 0 {
			entry.Errorf("パラメータ'%s'がありません", param)
			switch param {
			case "name":
				return errors.ParamNameMissingError()
			case "price":
				return errors.ParamPriceMissingError()
			}
		}
		entry.Infof("パラメータ'%s'が存在することを確認しました", param)
	}

	// パラメータNameの値が文字列でない場合
	if _, ok := interface{}(b.Name).(string); !ok {
		entry.Errorf("本の名前が文字列ではありません")
		return errors.BookNameNotStringError()
	}
	entry.Infof("パラメータ'name'が文字列であることを確認しました")

	// パラメータPriceの値が整数型でない場合
	if _, ok := interface{}(b.Price).(int); !ok {
		entry.Errorf("本の値段が整数型ではありません")
		return errors.BookPriceNotIntegerError()
	}
	entry.Infof("パラメータ'price'が整数型であることを確認しました")

	// bのパラメータにNameは存在するが空の場合
	if b.Name == "" {
		entry.Errorf("本の名前が空です")
		return errors.BookNameEmptyError()
	}
	entry.Infof("本の名前が空でないことを確認しました")

	// bのパラメータにPriceが存在するが空の場合
	if b.Price == 0 {
		entry.Errorf("本の価格が0です")
		return errors.BookPriceZeroError()
	}
	entry.Infof("本の価格が0でないことを確認しました")

	// bのパラメータにNameは存在するが50文字以上の場合
	if len(b.Name) > 50 {
		entry.Errorf("本の名前が長すぎます")
		return errors.BookNameTooLongError()
	}
	entry.Infof("本の名前が50文字以内であることを確認しました")

	// bのパラメータにPriceが存在するが20000より大きい場合
	if b.Price > 20000 {
		entry.Errorf("本の価格が高すぎます")
		return errors.BookPriceTooHighError()
	}
	entry.Infof("本の価格が20000円以下であることを確認しました")

	return nil
}

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

	b.ID = strconv.FormatInt(lastInsertId, 10)
	b.CreatedAt = time.Now()

	entry.Infof("本の登録に成功しました")
	entry.Infof("CreateBook関数が終了しました")
	return nil
}

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
		var createdAt string
		err := rows.Scan(&book.ID, &book.Name, &book.Price, &createdAt)
		if err != nil {
			entry.Errorf("データベース結果のスキャンに失敗しました: %v", err)
			return nil, errors.DatabaseScanError()
		}

		// 文字列からtime.Timeへの変換
		book.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			entry.Errorf("作成日時の変換に失敗しました: %v", err)
			return nil, errors.DatabaseScanError()
		}

		books = append(books, book)
	}
	entry.Infof("データベース結果のスキャンに成功しました")

	entry.Infof("GetBooks関数が終了しました")
	return books, nil
}
