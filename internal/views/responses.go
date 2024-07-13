package views

import (
	"go-api-tutorial/internal/errors"
	"go-api-tutorial/internal/transaction"
)

type Response struct {
	TrnID   string                 `json:"trn_id"`
	TrnTime string                 `json:"trn_time"`
	Result  map[string]interface{} `json:"result"`
}

func NewResponse(result map[string]interface{}) *Response {
	trnID, trnTime := transaction.InitializeTransaction()
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

func NewExceptionResponse(exception *errors.UserDefinedError) *ExceptionResponse {
	trnID, trnTime := transaction.InitializeTransaction()
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
	response["result"].(map[string]interface{})["error_info"] = e.ErrorInfo
	return response
}
