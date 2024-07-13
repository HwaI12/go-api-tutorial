package errors

import (
	"fmt"
	"net/http"
)

// カスタムエラー型 UserDefinedError を定義し、それに対するメソッドを実装
type UserDefinedError struct {
	ErrorCode      string `json:"error_code"`
	ErrorMessage   string `json:"error_message"`
	HTTPStatusCode int    `json:"-"`
}

// Error メソッドは UserDefinedError をエラーメッセージとしてフォーマットする
func (e *UserDefinedError) Error() string {
	return fmt.Sprintf("[%d] [%s] %s", e.HTTPStatusCode, e.ErrorCode, e.ErrorMessage)
}

func UnexpectedError() *UserDefinedError {
	return &UserDefinedError{"GOTA-Z-000-00", "予測不能エラーです", http.StatusInternalServerError}
}

func DatabaseError() *UserDefinedError {
	return &UserDefinedError{"GOTA-X-001-00", "データベースエラーです。もう一度お試しください。", http.StatusInternalServerError}
}

func DatabaseQueryError() *UserDefinedError {
	return &UserDefinedError{"GOTA-X-001-01", "データベースクエリの実行に失敗しました", http.StatusInternalServerError}
}

func DatabaseScanError() *UserDefinedError {
	return &UserDefinedError{"GOTA-X-001-02", "データベース結果のスキャンに失敗しました", http.StatusInternalServerError}
}

func DatabaseCloseError() *UserDefinedError {
	return &UserDefinedError{"GOTA-X-001-03", "データベース結果のクローズに失敗しました", http.StatusInternalServerError}
}

func SQLPreparationError() *UserDefinedError {
	return &UserDefinedError{"GOTA-X-001-04", "SQLステートメントの準備に失敗しました", http.StatusInternalServerError}
}

func DatabaseInsertError() *UserDefinedError {
	return &UserDefinedError{"GOTA-X-001-05", "データベースへの挿入に失敗しました", http.StatusInternalServerError}
}

func LastInsertIDError() *UserDefinedError {
	return &UserDefinedError{"GOTA-X-001-06", "最後に挿入されたIDの取得に失敗しました", http.StatusInternalServerError}
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

func APIKeyEmptyError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-021-00", "APIキーが空です。", http.StatusUnauthorized}
}

func InvalidAPIKeyError() *UserDefinedError {
	return &UserDefinedError{"GOTA-W-021-01", "APIキーが無効です。", http.StatusUnauthorized}
}
