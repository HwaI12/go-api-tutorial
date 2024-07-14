package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	error "github.com/HwaI12/go-api-tutorial/internal/error"
	logger "github.com/HwaI12/go-api-tutorial/internal/log"
	model "github.com/HwaI12/go-api-tutorial/internal/model"
	view "github.com/HwaI12/go-api-tutorial/internal/view"
)

// 書籍データに関する操作を行うコントローラー
type BookController struct {
	DB *sql.DB
}

// 新しい BookController を作成して返す
func NewBookController(db *sql.DB) *BookController {
	return &BookController{DB: db}
}

// 新しい書籍データをデータベースに登録するハンドラー
func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entry := logger.WithTransaction(ctx)

	var input model.BookInput
	entry.Infof("リクエストボディのデコードを開始します")
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		entry.Errorf("リクエストボディのデコードに失敗しました: %v", err)
		view.RespondWithError(w, ctx, error.InvalidRequestError())
		return
	}
	entry.Infof("リクエストボディのデコードに成功しました")

	// 入力データの検証
	book := model.Book{
		Name:  input.Name,
		Price: input.Price,
	}

	entry.Debugf("入力されたデータ: %+v", map[string]interface{}{
		"name":  book.Name,
		"price": book.Price,
	})

	entry.Infof("バリデーションを開始します")
	if err := book.Validate(ctx); err != nil {
		entry.Errorf("バリデーションに失敗しました: %v", err)
		view.RespondWithError(w, ctx, err.(*error.UserDefinedError))
		return
	}
	entry.Infof("バリデーションに成功しました")

	entry.Infof("本の登録を開始します")
	if err := book.CreateBook(ctx, c.DB); err != nil {
		entry.Errorf("本の登録に失敗しました: %v", err)
		view.RespondWithError(w, ctx, err.(*error.UserDefinedError))
		return
	}
	entry.Infof("本の登録に成功しました")

	entry.Infof("レスポンスを返却します")
	responseData := map[string]interface{}{
		"name":  book.Name,
		"price": book.Price,
	}
	response := view.CreateResponse(ctx, responseData)
	entry.Debugf("レスポンス結果: %+v", response)
	fmt.Printf("レスポンス結果: %+v", response)
	view.RespondWithJSON(w, ctx, http.StatusCreated, responseData)
	entry.Infof("レスポンスの返却に成功しました")
}

// データベースから書籍データを取得して返すハンドラー
func (c *BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entry := logger.WithTransaction(ctx)

	entry.Infof("本の一覧取得を開始します")
	books, err := model.GetBooks(ctx, c.DB)
	if err != nil {
		entry.Errorf("本の一覧取得に失敗しました: %v", err)
		view.RespondWithError(w, ctx, err.(*error.UserDefinedError))
		return
	}

	if len(books) == 0 {
		entry.Warnf("取得するデータがありません")
		view.RespondWithError(w, ctx, error.NoDataFoundError())
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
	response := view.CreateResponse(ctx, responseData)
	entry.Debugf("レスポンス結果: %+v", response)
	fmt.Printf("レスポンス結果: %+v", response)
	view.RespondWithJSON(w, ctx, http.StatusOK, responseData)
	entry.Infof("レスポンスの返却に成功しました")
}
