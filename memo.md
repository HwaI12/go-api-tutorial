# 学習メモ
## ディレクトリ構成
```sh
.
├── api
│   └── routes.go
├── cmd
│   └── myapp
│       └── main.go
├── internal
│   ├── controllers
│   │   └── book_controller.go
│   ├── models
│   │   └── book.go
│   ├── views
│   │   └── responses.go
│   └── errors
│       └── custom_errors.go
├── go.mod
├── go.sum
├── .env
└── pkg
    └── database
        └── database.go
```
- cmd/
  - アプリのエントリーポイント
  - ここではAPIのエントリーポイントを配置
- internal/
  - 内部パッケージ。外部には公開しない。
    - controllers/
      - コントローラーのロジック
    - models/
      - データモデル。
    - views/
      - レスポンスのビュー。
- pkg/
  - 再利用可能なパッケージ
  - ここではデータベース接続ロジックを配置
- api/: ルーティングの設定

## MVCモデルとは
- Models
  - 役割: データベースとのやり取りを管理し、データの構造を定義する
  - 例: ユーザーの情報をデータベースに保存する、タスクの情報を取得する
- Controllers
  - 役割: HTTPリクエストを受け取り、適切なモデルを呼び出してデータベース操作を行い、レスポンスを返す
  - 例: ユーザーが新規登録するときの処理を行う、タスクを作成するリクエストを処理する
- View
  - 役割: ユーザーに表示するためのデータを整形する、WebアプリケーションではHTMLやJSONレスポンスなどを生成する
  - 例: ユーザー情報をJSON形式で返す、タスク一覧をHTML形式で表示する

## ListenAndServe
ListenAndServeとは、指定したアドレスとポートでHTTPサーバを起動するメソッド

## カスタムエラー処理
### エラーコード
|エラーコード|ステータスコード|エラーメッセージ|
|---|---|---|
|GOTA-Z-000-00|500|予測不能エラーです|
|GOTA-X-001-00|500|データベースエラーです。もう一度お試しください。|
|GOTA-W-011-00|400|パラメータ'name'がありません。|
|GOTA-W-011-01|400|パラメータ'price'がありません。|
|GOTA-W-011-02|400|本の名前がありません。|
|GOTA-W-011-03|400|本の値段がありません。|
|GOTA-W-011-04|400|本の名前が文字列ではありません。|
|GOTA-W-011-05|400|本の値段が整数型ではありません。|
|GOTA-W-011-06|400|本の名前が空です。1文字以上書いてください。|
|GOTA-W-011-07|400|本の値段が空です。1文字以上書いてください。|
|GOTA-W-011-08|400|本の名前が長すぎます。50文字以内で書いてください。|
|GOTA-W-011-09|400|本の値段が高すぎます。20000円以内で書いてください。|
|GOTA-W-021-00|401|APIキーが無効です。|

### エラーコードレベル
|レベルコード|状態|
|---|---|
|C|トレース|
|D|デバッグ|
|I|情報|
|W|警告|
|X|エラー|
|Z|致命的|

### 分類コードレベル
|分類コード|分類|
|---|---|
|000|重大エラー|
|001|システム例外エラー|
|011|バリデーションエラー|
|021|認証エラー|
|031|URLエラー|

### エラーメッセージ
#### パラメータ存在チェック
- パラメータ'name'がありません。
- パラメータ'price'がありません。

#### 中身検証チェック
##### Noneの場合
- 本の名前がありません。
- 本の値段がありません。

##### 型
- 本の名前が文字列型ではありません。文字列型にしてください。
- 本の値段が整数型ではありません。整数型にしてください。

##### 空で渡されてる場合
- 本の名前が空です。1文字以上書いてください。
- 本の値段が空です。1文字以上書いてください。

##### 制限
- 本の名前が重複しています。別の名前を入力してください。
- 本の名前が長すぎます。50文字以内で書いてください。
- 本の値段がマイナスです。正の値を入力してください。
- 本の値段が高すぎます。20000円以内で書いてください。

##### データベースエラー
- 本を取得できませんでした。
- 本を登録できませんでした。

#### 予期しないエラー
- 予測不能エラーです。

#### APIキー
- APIキーが無効です。

#### データベースと連携
- データベースエラーです。もう一度お試しください。


### エラーの定義
`internal/errors/custom_errors.go`に記載

今の書き方と以下の書き方どっちがいいんだろう
```go
var (
	UnexpectedError          = &UserDefinedError{"GOTA-Z-000-00", "予測不能エラーです", http.StatusInternalServerError}
	DatabaseError            = &UserDefinedError{"GOTA-X-001-00", "データベースエラーです。もう一度お試しください。", http.StatusInternalServerError}
	ParamNameMissingError    = &UserDefinedError{"GOTA-W-011-00", "パラメータ'name'がありません。", http.StatusBadRequest}
	ParamPriceMissingError   = &UserDefinedError{"GOTA-W-011-01", "パラメータ'price'がありません。", http.StatusBadRequest}
	BookNameMissingError     = &UserDefinedError{"GOTA-W-011-02", "本の名前がありません。", http.StatusBadRequest}
	BookPriceMissingError    = &UserDefinedError{"GOTA-W-011-03", "本の値段がありません。", http.StatusBadRequest}
	BookNameNotStringError   = &UserDefinedError{"GOTA-W-011-04", "本の名前が文字列ではありません。", http.StatusBadRequest}
	BookPriceNotIntegerError = &UserDefinedError{"GOTA-W-011-05", "本の値段が整数型ではありません。", http.StatusBadRequest}
	BookNameEmptyError       = &UserDefinedError{"GOTA-W-011-06", "本の名前が空です。1文字以上書いてください。", http.StatusBadRequest}
	BookPriceEmptyError      = &UserDefinedError{"GOTA-W-011-07", "本の値段が空です。1文字以上書いてください。", http.StatusBadRequest}
	BookNameTooLongError     = &UserDefinedError{"GOTA-W-011-08", "本の名前が長すぎます。50文字以内で書いてください。", http.StatusBadRequest}
	BookPriceTooHighError    = &UserDefinedError{"GOTA-W-011-09", "本の値段が高すぎます。20000円以内で書いてください。", http.StatusBadRequest}
	InvalidAPIKeyError       = &UserDefinedError{"GOTA-W-021-00", "APIキーが無効です。", http.StatusUnauthorized}
)

// NewCustomError()は、UserDefinedError型の関数
// この関数は、エラーコード、エラーメッセージ、HTTPステータスコードを指定して、カスタムエラーを作成する
func NewCustomError(errorCode, errorMessage string, httpStatusCode int) *UserDefinedError {
	return &UserDefinedError{
		ErrorCode:      errorCode,
		ErrorMessage:   errorMessage,
		HTTPStatusCode: httpStatusCode,
	}
}
```

## MySQL
- MySQLのインストール
  ```sh
  brew install mysql
  ```
- MySQLの起動
  ```sh
  brew services start mysql
  ```
- MySQLの設定
  ```sh
  mysql -u root -p
  ```
- データベースのテーブルを作成
  ```sql
  USE book_db;
  CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    price INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  );
  ```