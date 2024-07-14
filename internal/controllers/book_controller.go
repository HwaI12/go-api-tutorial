package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/HwaI12/go-api-tutorial/internal/errors"
	"github.com/HwaI12/go-api-tutorial/internal/logger"
	"github.com/HwaI12/go-api-tutorial/internal/models"
	"github.com/HwaI12/go-api-tutorial/internal/views"
	"github.com/sirupsen/logrus"
)

// BookController は書籍データに関する操作を行うコントローラーである
type BookController struct {
	DB *sql.DB
}

// NewBookController は新しい BookController を作成して返す
func NewBookController(db *sql.DB) *BookController {
	return &BookController{DB: db}
}

// CreateBook は新しい書籍データをデータベースに登録するハンドラーである
func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entry := logger.WithTransaction(ctx)
	var book models.Book

	entry.Infof("リクエストボディのデコードを開始します")
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		handleError(w, ctx, entry, errors.InvalidRequestError(), "リクエストボディのデコードに失敗しました: %v", err)
		return
	}
	entry.Infof("リクエストボディのデコードに成功しました")

	entry.Debugf("入力されたデータ: %+v", map[string]interface{}{
		"name":  book.Name,
		"price": book.Price,
	})

	entry.Infof("バリデーションを開始します")
	if err := book.Validate(ctx); err != nil {
		handleError(w, ctx, entry, err.(*errors.UserDefinedError), "バリデーションに失敗しました: %v", err)
		return
	}
	entry.Infof("バリデーションに成功しました")

	entry.Infof("本の登録を開始します")
	if err := book.CreateBook(ctx, c.DB); err != nil {
		handleError(w, ctx, entry, err.(*errors.UserDefinedError), "本の登録に失敗しました: %v", err)
		return
	}
	entry.Infof("本の登録に成功しました")

	entry.Infof("レスポンスを返却します")
	responseData := map[string]interface{}{
		"name":  book.Name,
		"price": book.Price,
	}
	response := views.CreateResponse(ctx, responseData)
	entry.Debugf("レスポンス結果: %+v", response)
	views.RespondWithJSON(w, ctx, http.StatusCreated, responseData)
	entry.Infof("レスポンスの返却に成功しました")
}

// GetBooks はデータベースから書籍データを取得して返すハンドラーである
func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entry := logger.WithTransaction(ctx)

	entry.Infof("本の一覧取得を開始します")
	books, err := models.GetBooks(ctx, c.DB)
	if err != nil {
		handleError(w, ctx, entry, err.(*errors.UserDefinedError), "本の一覧取得に失敗しました: %v", err)
		return
	}

	if len(books) == 0 {
		handleError(w, ctx, entry, errors.NoDataFoundError(), "本の一覧が空です")
		return
	}
	entry.Infof("本の一覧取得に成功しました")

	// データを変換する
	bookList := make([]map[string]interface{}, len(books))
	for i, book := range books {
		bookList[i] = map[string]interface{}{
			"id":         book.ID,
			"name":       book.Name,
			"price":      book.Price,
			"created_at": book.CreatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	entry.Infof("レスポンスを返却します")
	responseData := map[string]interface{}{
		"books": bookList,
	}
	response := views.CreateResponse(ctx, responseData)
	entry.Debugf("レスポンス結果: %+v", response)
	views.RespondWithJSON(w, ctx, http.StatusOK, responseData)
	entry.Infof("レスポンスの返却に成功しました")
}

// handleError はエラーレスポンスを返す共通のハンドリング関数である
func handleError(w http.ResponseWriter, ctx context.Context, entry *logrus.Entry, err *errors.UserDefinedError, message string, args ...interface{}) {
	if len(args) > 0 {
		entry.Errorf(message, args...)
	} else {
		entry.Errorf(message)
	}
	entry.Infof("レスポンスを返却します")
	entry.Debugf("%+v", err)
	response := views.CreateExceptionResponse(ctx, err)
	entry.Debugf("レスポンス結果: %+v", response)
	views.RespondWithError(w, ctx, err)
	entry.Infof("レスポンスの返却に成功しました")
}
