package errors

import (
	"fmt"
	"net/http"
)

// UserDefinedError カスタムエラー型
type UserDefinedError struct {
	ErrorCode      string `json:"error_code"`
	ErrorMessage   string `json:"error_message"`
	HTTPStatusCode int    `json:"-"`
}

// Error メソッドは UserDefinedError をエラーメッセージとしてフォーマットする
func (e *UserDefinedError) Error() string {
	return fmt.Sprintf("[%d] [%s] %s", e.HTTPStatusCode, e.ErrorCode, e.ErrorMessage)
}

// エラー生成関数

func UnexpectedError() *UserDefinedError {
	return &UserDefinedError{"BUSN-ERR-500-00", "予測不能エラーです", http.StatusInternalServerError}
}

func EnvLoadError() *UserDefinedError {
	return &UserDefinedError{"ENV-ERR-500-00", ".envファイルの読み込みに失敗しました", http.StatusInternalServerError}
}

func DatabaseConnectionError() *UserDefinedError {
	return &UserDefinedError{"DB-ERR-500-00", "データベースへの接続に失敗しました", http.StatusInternalServerError}
}

func DatabaseQueryError() *UserDefinedError {
	return &UserDefinedError{"DB-ERR-500-01", "データベースクエリの実行に失敗しました", http.StatusInternalServerError}
}

func DatabaseScanError() *UserDefinedError {
	return &UserDefinedError{"DB-ERR-500-02", "データベース結果のスキャンに失敗しました", http.StatusInternalServerError}
}

func DatabaseCloseError() *UserDefinedError {
	return &UserDefinedError{"DB-ERR-500-03", "データベース結果のクローズに失敗しました", http.StatusInternalServerError}
}

func SQLPreparationError() *UserDefinedError {
	return &UserDefinedError{"DB-ERR-500-04", "SQLステートメントの準備に失敗しました", http.StatusInternalServerError}
}

func DatabaseInsertError() *UserDefinedError {
	return &UserDefinedError{"DB-ERR-500-05", "データベースへの挿入に失敗しました", http.StatusInternalServerError}
}

func LastInsertIDError() *UserDefinedError {
	return &UserDefinedError{"DB-ERR-500-06", "最後に挿入されたIDの取得に失敗しました", http.StatusInternalServerError}
}

func DatabaseSelectError() *UserDefinedError {
	return &UserDefinedError{"DB-ERR-500-07", "データベースからの取得に失敗しました", http.StatusInternalServerError}
}

func NoDataFoundError() *UserDefinedError {
	return &UserDefinedError{"DB-ERR-404-00", "取得するデータがありません", http.StatusNotFound}
}

func ServerStartError() *UserDefinedError {
	return &UserDefinedError{"SRV-ERR-500-00", "サーバーの起動に失敗しました", http.StatusInternalServerError}
}

func ServerShutdownError() *UserDefinedError {
	return &UserDefinedError{"SRV-ERR-500-01", "サーバーのシャットダウンに失敗しました", http.StatusInternalServerError}
}

func APIKeyEmptyError() *UserDefinedError {
	return &UserDefinedError{"AUTH-ERR-401-00", "APIキーが空です", http.StatusUnauthorized}
}

func InvalidAPIKeyError() *UserDefinedError {
	return &UserDefinedError{"AUTH-ERR-401-01", "APIキーが無効です", http.StatusUnauthorized}
}

func InvalidRequestError() *UserDefinedError {
	return &UserDefinedError{"VAL-ERR-400-07", "リクエストボディのデコードに失敗しました", http.StatusBadRequest}
}

func ParamNameMissingError() *UserDefinedError {
	return &UserDefinedError{"VAL-ERR-400-00", "パラメータ'name'がありません。パラメータを正しく設定するか、値を入力してください", http.StatusBadRequest}
}

func ParamPriceMissingError() *UserDefinedError {
	return &UserDefinedError{"VAL-ERR-400-01", "パラメータ'price'がありません。パラメータを正しく設定するか、値を入力してください", http.StatusBadRequest}
}

func BookNameEmptyError() *UserDefinedError {
	return &UserDefinedError{"VAL-ERR-400-02", "パラメータ'name'が空です。本の名前を入力してください", http.StatusBadRequest}
}

func BookPriceEmptyError() *UserDefinedError {
	return &UserDefinedError{"VAL-ERR-400-03", "パラメータ'price'が0です。本の価格を入力してください", http.StatusBadRequest}
}

func BookNameTooLongError() *UserDefinedError {
	return &UserDefinedError{"VAL-ERR-400-04", "パラメータ'name'が長すぎます。50文字以内で書いてください", http.StatusBadRequest}
}

func BookPriceNegativeError() *UserDefinedError {
	return &UserDefinedError{"VAL-ERR-400-05", "パラメータ'price'が0以下です。正の整数を入力してください", http.StatusBadRequest}
}

func BookPriceTooHighError() *UserDefinedError {
	return &UserDefinedError{"VAL-ERR-400-06", "パラメータ'price'が高すぎます。20000円以内で書いてください", http.StatusBadRequest}
}
