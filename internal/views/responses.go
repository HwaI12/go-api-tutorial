package views

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/HwaI12/go-api-tutorial/internal/errors"
	"github.com/HwaI12/go-api-tutorial/internal/transaction"
)

type Response struct {
	TrnID   string                 `json:"trn_id"`
	TrnTime string                 `json:"trn_time"`
	Result  map[string]interface{} `json:"result"`
}

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

func (r *Response) Build() map[string]interface{} {
	response := map[string]interface{}{
		"trn_id":   r.TrnID,
		"trn_time": r.TrnTime,
		"result":   r.Result,
	}
	return response
}

type ExceptionResponse struct {
	Response
	ErrorInfo map[string]string `json:"error_info"`
}

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

func (e *ExceptionResponse) Build() map[string]interface{} {
	response := e.Response.Build()
	response["error_info"] = e.ErrorInfo
	return response
}

func RespondWithError(w http.ResponseWriter, ctx context.Context, err *errors.UserDefinedError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPStatusCode)
	response := CreateExceptionResponse(ctx, err)
	json.NewEncoder(w).Encode(response.Build())
}

func RespondWithJSON(w http.ResponseWriter, ctx context.Context, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := CreateResponse(ctx, map[string]interface{}{
		"payload": payload,
	})
	json.NewEncoder(w).Encode(response.Build())
}
