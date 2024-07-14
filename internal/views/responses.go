package views

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HwaI12/go-api-tutorial/internal/errors"
	"github.com/HwaI12/go-api-tutorial/internal/transaction"
)

type Response struct {
	TrnID   string      `json:"trn_id"`
	TrnTime string      `json:"trn_time"`
	Result  interface{} `json:"result"`
}

// 正常なレスポンスを作成して返す
func CreateResponse(ctx context.Context, result interface{}) *Response {
	trnID := ctx.Value(transaction.TrnIDKey).(string)
	trnTime := ctx.Value(transaction.TrnTimeKey).(string)
	return &Response{
		TrnID:   trnID,
		TrnTime: trnTime,
		Result:  result,
	}
}

// エラーレスポンスを作成するための構造体
type ExceptionResponse struct {
	TrnID   string `json:"trn_id"`
	TrnTime string `json:"trn_time"`
	Result  struct {
		ErrorCode    string `json:"error_code"`
		ErrorMessage string `json:"error_message"`
	} `json:"result"`
}

// エラーレスポンスを作成して返す
func CreateExceptionResponse(ctx context.Context, exception *errors.UserDefinedError) *ExceptionResponse {
	trnID := ctx.Value(transaction.TrnIDKey).(string)
	trnTime := ctx.Value(transaction.TrnTimeKey).(string)
	return &ExceptionResponse{
		TrnID:   trnID,
		TrnTime: trnTime,
		Result: struct {
			ErrorCode    string `json:"error_code"`
			ErrorMessage string `json:"error_message"`
		}{
			ErrorCode:    exception.ErrorCode,
			ErrorMessage: exception.ErrorMessage,
		},
	}
}

// エラーレスポンスを返す
// エラーレスポンスはエラーコードとエラーメッセージを含む
// エラーコードとエラーメッセージはユーザー定義エラーから取得
// エラーレスポンスはJSON形式で返す
func RespondWithError(w http.ResponseWriter, ctx context.Context, err *errors.UserDefinedError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPStatusCode)
	response := CreateExceptionResponse(ctx, err)
	json.NewEncoder(w).Encode(response)
}

// JSON形式のレスポンスを返す
// レスポンスは正常な場合とエラーの場合がある
// レスポンスはJSON形式で返す
func RespondWithJSON(w http.ResponseWriter, ctx context.Context, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := CreateResponse(ctx, payload)
	json.NewEncoder(w).Encode(response)
}
