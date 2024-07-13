- [学習メモ](#学習メモ)
  - [ディレクトリ構成](#ディレクトリ構成)
  - [MVCモデルとは](#mvcモデルとは)
  - [ListenAndServe](#listenandserve)
  - [カスタムエラー処理](#カスタムエラー処理)
    - [エラーコード](#エラーコード)
    - [エラーコードとメッセージ](#エラーコードとメッセージ)
    - [エラーコードレベル](#エラーコードレベル)
    - [分類コードレベル](#分類コードレベル)
    - [エラーメッセージ](#エラーメッセージ)
      - [パラメータ存在チェック](#パラメータ存在チェック)
      - [中身検証チェック](#中身検証チェック)
        - [Noneの場合](#noneの場合)
        - [型](#型)
        - [空で渡されてる場合](#空で渡されてる場合)
        - [制限](#制限)
        - [データベースエラー](#データベースエラー)
      - [予期しないエラー](#予期しないエラー)
      - [APIキー](#apiキー)
      - [データベースと連携](#データベースと連携)
    - [エラーの定義](#エラーの定義)
  - [MySQL](#mysql)
  - [ログの追加](#ログの追加)

# 学習メモ
## ディレクトリ構成
```sh
.
├── README.md
├── api
│   └── routes.go
├── cmd
│   └── myapp
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── Log
│   │   └── logger.go
│   ├── controllers
│   │   └── book_controller.go
│   ├── errors
│   │   └── custom_errors.go
│   ├── middleware
│   │   └── auth_middleware.go
│   ├── models
│   │   └── book.go
│   ├── transaction
│   │   └── transaction.go
│   └── views
│       └── responses.go
├── memo.md
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
データベースエラーのステータスコードは一般的に500番（内部サーバーエラー）を使用します。これはサーバー内部で予期しないエラーが発生した場合に使用されるステータスコードです。以下は、データベース関連のエラーに対応するステータスコードの一覧です。

### エラーコードとメッセージ

|エラーコード|ステータスコード|エラーメッセージ|
|---|---|---|
|GOTA-Z-000-00|500|予測不能エラーです|
|GOTA-X-001-00|500|データベースエラーです。もう一度お試しください。|
|GOTA-X-001-01|500|データベースクエリの実行に失敗しました|
|GOTA-X-001-02|500|データベース結果のスキャンに失敗しました|
|GOTA-X-001-03|500|データベース結果のクローズに失敗しました|
|GOTA-X-001-04|500|SQLステートメントの準備に失敗しました|
|GOTA-X-001-05|500|データベースへの挿入に失敗しました|
|GOTA-X-001-06|500|最後に挿入されたIDの取得に失敗しました|
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
|GOTA-W-021-00|401|APIキーが空です。|
|GOTA-W-021-01|401|APIキーが無効です。|

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
- APIキーが空です。
- APIキーが無効です。

#### データベースと連携
- データベースクエリの実行に失敗しました
- データベース結果のスキャンに失敗しました
- データベース結果のクローズに失敗しました
- SQLステートメントの準備に失敗しました
- データベースへの挿入に失敗しました
- 最後に挿入されたIDの取得に失敗しました


### エラーの定義
`internal/errors/custom_errors.go`に記載

```go
func (e *UserDefinedError) Error() string {
	return fmt.Sprintf("[%d] [%s] %s", e.HTTPStatusCode, e.ErrorCode, e.ErrorMessage)
}

func UnexpectedError() *UserDefinedError {
	return &UserDefinedError{"GOTA-Z-000-00", "予測不能エラーです", http.StatusInternalServerError}
}

func DatabaseError() *UserDefinedError {
	return &UserDefinedError{"GOTA-X-001-00", "データベースエラーです。もう一度お試しください。", http.StatusInternalServerError}
}

・・・
```

```go
func (e *UserDefinedError) Error() string {
	return fmt.Sprintf("[%d] [%s] %s", e.HTTPStatusCode, e.ErrorCode, e.ErrorMessage)
}

var (
	UnexpectedError          = &UserDefinedError{"GOTA-Z-000-00", "予測不能エラーです", http.StatusInternalServerError}
	DatabaseError            = &UserDefinedError{"GOTA-X-001-00", "データベースエラーです。もう一度お試しください。", http.StatusInternalServerError}

  ・・・
)
func NewCustomError(errorCode, errorMessage string, httpStatusCode int) *UserDefinedError {
	return &UserDefinedError{
		ErrorCode:      errorCode,
		ErrorMessage:   errorMessage,
		HTTPStatusCode: httpStatusCode,
	}
}
```

上の書き方と下の書き方、どっちがいいのだろうか？

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

## ログの追加
ログレベルを設定し別ファイルにログを保存するようにした。
![Log-image](https://github.com/user-attachments/assets/9a0a4aaa-d5d6-4fdd-98ec-316613cdf010)