# 学習メモ
## ディレクトリ構成
```sh
.
├── .env
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

## Go言語の関数やメソッドの書き方メモ
- ListenAndServe
  - 指定したアドレスとポートで HTTP サーバを起動するメソッド
- GoにはNoneがない！？
  - パラメータがあるかどうかのチェックを "" or 0 で判断する
  - パラメータの処理と・値の処理が同じ処理方法になる

## カスタム例外
| エラーコード    | HTTPステータスコード | エラーメッセージ                                                                  |
| --------------- | -------------------- | --------------------------------------------------------------------------------- |
| VAL-ERR-400-00  | 400                  | パラメータ'name'がありません。パラメータを正しく設定するか、値を入力してください  |
| VAL-ERR-400-01  | 400                  | パラメータ'price'がありません。パラメータを正しく設定するか、値を入力してください |
| VAL-ERR-400-02  | 400                  | パラメータ'name'が空です。本の名前を入力してください                              |
| VAL-ERR-400-03  | 400                  | パラメータ'price'が0です。本の価格を入力してください                              |
| VAL-ERR-400-04  | 400                  | パラメータ'name'が長すぎます。50文字以内で書いてください"                         |
| VAL-ERR-400-05  | 400                  | パラメータ'price'が0以下です。正の整数を入力してください                          |
| VAL-ERR-400-06  | 400                  | パラメータ'price'が高すぎます。20000円以内で書いてください                        |
| BUSN-ERR-500-00 | 500                  | 予測不能エラーです。                                                              |
| ENV-ERR-500-00  | 500                  | .envファイルの読み込みに失敗しました。                                            |
| DB-ERR-500-00   | 500                  | データベースへの接続に失敗しました。                                              |
| DB-ERR-500-01   | 500                  | データベースクエリの実行に失敗しました。                                          |
| DB-ERR-500-02   | 500                  | データベース結果のスキャンに失敗しました。                                        |
| DB-ERR-500-03   | 500                  | データベース結果のクローズに失敗しました。                                        |
| DB-ERR-500-04   | 500                  | SQLステートメントの準備に失敗しました。                                           |
| DB-ERR-500-05   | 500                  | データベースへの挿入に失敗しました。                                              |
| DB-ERR-500-06   | 500                  | 最後に挿入されたIDの取得に失敗しました。                                          |
| DB-ERR-500-07   | 500                  | データベースからの取得に失敗しました。                                            |
| DB-ERR-404-00   | 404                  | 取得するデータがありません。                                                      |
| SRV-ERR-500-00  | 500                  | サーバーの起動に失敗しました。                                                    |
| SRV-ERR-500-01  | 500                  | サーバーのシャットダウンに失敗しました。                                          |
| AUTH-ERR-401-00 | 401                  | APIキーが空です。                                                                 |
| AUTH-ERR-401-01 | 401                  | APIキーが無効です。                                                               |
| VAL-ERR-400-07  | 400                  | リクエストボディのデコードに失敗しました。                                        |

### 設定した規則
- **BUSN-ERR-500-00**: ビジネスロジックで発生する予測不能なエラー。
- **DB-ERR-500-xx**: データベースに関連するエラー。
- **VAL-ERR-ERR-400-xx**: バリデーションエラー（ユーザー入力の検証に失敗した場合）。
- **AUTH-ERR-401-xx**: 認証エラー。

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
Dockerで構築していない + 簡易的なアプリなため自力で作成する必要がある
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
- 全てのデータの削除
  ```sql
  truncate table books;
  ```

## ログの追加
ログレベルを設定し別ファイルにログを保存するようにした。
![Log-image](https://github.com/user-attachments/assets/9a0a4aaa-d5d6-4fdd-98ec-316613cdf010)