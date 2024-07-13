package errors

import (
	"fmt"
	"net/http"
)

// typeとは、型を定義するもの
// この場合、UserDefinedErrorという型を定義している
type UserDefinedError struct {
	ErrorCode      string `json:"error_code"`
	ErrorMessage   string `json:"error_message"`
	HTTPStatusCode int    `json:"-"`
}

// Error()は、UserDefinedError型のメソッド
func (e *UserDefinedError) Error() string {
	// fmt.Sprintf()は、文字列をフォーマットする関数
	// この場合、HTTPステータスコード、エラーコード、エラーメッセージを文字列として返す
	return fmt.Sprintf("[%d] [%s] %s", e.HTTPStatusCode, e.ErrorCode, e.ErrorMessage)
}

func UnexpectedError() *UserDefinedError {
	return &UserDefinedError{"GOTA-Z-000-00", "予測不能エラーです", http.StatusInternalServerError}
}

func DatabaseError() *UserDefinedError {
	return &UserDefinedError{"GOTA-X-001-00", "データベースエラーです。もう一度お試しください。", http.StatusInternalServerError}
}

func ParamNameMissingError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-011-00", "パラメータ'name'がありません。", http.StatusBadRequest}
}

func ParamPriceMissingError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-011-01", "パラメータ'price'がありません。", http.StatusBadRequest}
}

func BookNameMissingError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-011-02", "本の名前がありません。", http.StatusBadRequest}
}

func BookPriceMissingError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-011-03", "本の値段がありません。", http.StatusBadRequest}
}

func BookNameNotStringError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-011-04", "本の名前が文字列ではありません。", http.StatusBadRequest}
}

func BookPriceNotIntegerError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-011-05", "本の値段が整数型ではありません。", http.StatusBadRequest}
}

func BookNameEmptyError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-011-06", "本の名前が空です。1文字以上書いてください。", http.StatusBadRequest}
}

func BookPriceEmptyError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-011-07", "本の値段が空です。1文字以上書いてください。", http.StatusBadRequest}
}

func BookNameTooLongError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-011-08", "本の名前が長すぎます。50文字以内で書いてください。", http.StatusBadRequest}
}

func BookPriceTooHighError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-011-09", "本の値段が高すぎます。20000円以内で書いてください。", http.StatusBadRequest}
}

func InvalidAPIKeyError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-021-00", "APIキーが無効です。", http.StatusUnauthorized}
}
