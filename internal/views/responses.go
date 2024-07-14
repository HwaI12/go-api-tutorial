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

type JSONResponse struct {
	TrnID   string                 `json:"trn_id"`
	TrnTime string                 `json:"trn_time"`
	Result  map[string]interface{} `json:"result"`
}

type JSONErrorResponse struct {
	TrnID   string            `json:"trn_id"`
	TrnTime string            `json:"trn_time"`
	Result  map[string]string `json:"result"`
}

func CreateResponse(ctx context.Context, result map[string]interface{}) *Response {
	trnID := ctx.Value(transaction.TrnIDKey).(string)
	trnTime := ctx.Value(transaction.TrnTimeKey).(string)
	return &Response{
		TrnID:   trnID,
		TrnTime: trnTime,
		Result:  result,
	}
}

func (r *Response) Build() JSONResponse {
	return JSONResponse{
		TrnID:   r.TrnID,
		TrnTime: r.TrnTime,
		Result:  r.Result,
	}
}

type ExceptionResponse struct {
	Response
	ErrorInfo map[string]string `json:"result"`
}

func CreateExceptionResponse(ctx context.Context, exception *errors.UserDefinedError) *ExceptionResponse {
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

func (e *ExceptionResponse) Build() JSONErrorResponse {
	return JSONErrorResponse{
		TrnID:   e.TrnID,
		TrnTime: e.TrnTime,
		Result:  e.ErrorInfo,
	}
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
	response := CreateResponse(ctx, payload.(map[string]interface{}))
	json.NewEncoder(w).Encode(response.Build())
}
