package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/HwaI12/go-api-tutorial/internal/logger"
	"github.com/HwaI12/go-api-tutorial/internal/models"
)

type BookController struct {
	DB *sql.DB
}

func NewBookController(db *sql.DB) *BookController {
	return &BookController{DB: db}
}

func (bc *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entry := logger.WithTransaction(ctx)

	entry.Info("リクエストを受信しました")

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		entry.WithError(err).Error("リクエストボディのデコードに失敗しました")
		http.Error(w, "無効なリクエストボディです", http.StatusBadRequest)
		return
	}

	entry.Debug("リクエストボディのデコードに成功しました")

	if err := book.Validate(); err != nil {
		entry.WithError(err).Warn("バリデーションエラー")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entry.Debug("バリデーションに成功しました")

	if err := book.CreateBook(bc.DB); err != nil {
		entry.WithError(err).Error("本の作成に失敗しました")
		http.Error(w, "本の作成に失敗しました", http.StatusInternalServerError)
		return
	}

	entry.Info("本の作成に成功しました")

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(book); err != nil {
		entry.WithError(err).Error("レスポンスのエンコードに失敗しました")
	}
	entry.Debug("レスポンスを送信しました")
}
