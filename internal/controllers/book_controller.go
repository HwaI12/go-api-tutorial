package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/HwaI12/go-api-tutorial/internal/errors"
	"github.com/HwaI12/go-api-tutorial/internal/logger"
	"github.com/HwaI12/go-api-tutorial/internal/models"
	"github.com/HwaI12/go-api-tutorial/internal/views"
)

type BookController struct {
	DB *sql.DB
}

func NewBookController(db *sql.DB) *BookController {
	return &BookController{DB: db}
}

func (c *BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entry := logger.WithTransaction(ctx)

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		entry.Errorf("リクエストボディのデコードに失敗しました: %v", err)
		views.RespondWithError(w, ctx, errors.InvalidRequestError())
		return
	}

	if err := book.Validate(ctx); err != nil {
		entry.Errorf("バリデーションに失敗しました: %v", err)
		views.RespondWithError(w, ctx, err.(*errors.UserDefinedError))
		return
	}

	if err := book.CreateBook(ctx, c.DB); err != nil {
		entry.Errorf("本の登録に失敗しました: %v", err)
		views.RespondWithError(w, ctx, err.(*errors.UserDefinedError))
		return
	}

	entry.Info("本の登録に成功しました")
	views.RespondWithJSON(w, ctx, http.StatusCreated, map[string]interface{}{
		"name":  book.Name,
		"price": book.Price,
	})
}
