package views

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HwaI12/go-api-tutorial/internal/errors"
	"github.com/HwaI12/go-api-tutorial/internal/transaction"
)

// トランザクションID、トランザクション時間、および結果を含むAPIレスポンスを表すResponse構造体
type Response struct {
	TrnID   string                 `json:"trn_id"`
	TrnTime string                 `json:"trn_time"`
	Result  map[string]interface{} `json:"result"`
}

// CreateResponse 関数は、新しいトランザクションを初期化し、その情報を含む Response を生成する。
func CreateResponse(ctx context.Context, result map[string]interface{}) *Response {
	ctx = transaction.InitializeTransaction(ctx)
	trnID := ctx.Value(transaction.TrnIDKey).(string)
	trnTime := ctx.Value(transaction.TrnTimeKey).(string)
	return &Response{
		TrnID:   trnID,
		TrnTime: trnTime,
		Result:  result,
	}
}

// Build メソッドは、Response 構造体を map[string]interface{} 型のレスポンスに変換する。
func (r *Response) Build() map[string]interface{} {
	response := map[string]interface{}{
		"trn_id":   r.TrnID,
		"trn_time": r.TrnTime,
		"result":   r.Result,
	}
	return response
}

// ExceptionResponse 構造体は、エラーレスポンスを表し、Response 構造体に加えてエラー情報を含む。
type ExceptionResponse struct {
	Response
	ErrorInfo map[string]string `json:"error_info"`
}

// CreateExceptionResponse 関数は、新しいトランザクションを初期化し、エラー情報を含む ExceptionResponse を生成する。
func CreateExceptionResponse(ctx context.Context, exception *errors.UserDefinedError) *ExceptionResponse {
	ctx = transaction.InitializeTransaction(ctx)
	trnID := ctx.Value(transaction.TrnIDKey).(string)
	trnTime := ctx.Value(transaction.TrnTimeKey).(string)
	return &ExceptionResponse{
		Response: Response{
			TrnID:   trnID,
			TrnTime: trnTime,
			Result:  map[string]interface{}{},
		},
		ErrorInfo: map[string]string{
			"error_code":    exception.ErrorCode,
			"error_message": exception.ErrorMessage,
		},
	}
}

// Build メソッドは、ExceptionResponse 構造体を map[string]interface{} 型のレスポンスに変換する。
func (e *ExceptionResponse) Build() map[string]interface{} {
	response := e.Response.Build()
	response["error_info"] = e.ErrorInfo
	return response
}

// RespondWithError 関数は、エラー情報を含むHTTPレスポンスをクライアントに返す。
func RespondWithError(w http.ResponseWriter, ctx context.Context, err *errors.UserDefinedError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPStatusCode)
	response := CreateExceptionResponse(ctx, err)
	json.NewEncoder(w).Encode(response.Build())
}

// RespondWithJSON 関数は、成功した結果を含むHTTPレスポンスをクライアントに返す。
func RespondWithJSON(w http.ResponseWriter, ctx context.Context, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := CreateResponse(ctx, map[string]interface{}{
		"payload": payload,
	})
	json.NewEncoder(w).Encode(response.Build())
}
