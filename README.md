# go-api-tutorial

簡単にGo言語とMySQLで欲しい本の名前や値段などの情報を管理するAPIを作成してみた！
## メモ
[memo.md](https://github.com/HwaI12/go-api-tutorial/blob/main/memo.md)

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
## .envファイル
```.env
DB_USER=root # データベースユーザー名
DB_PASSWORD=sample123 # データベースパスワード
DB_NAME=book_db # データベース名
DB_HOST=localhost # データベースホスト名またはIPアドレス
DB_PORT=3306 # データベースポート番号
API_KEY=123456789 # APIキー
```