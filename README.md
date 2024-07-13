# go-api-tutorial

簡単にGo言語とMySQLで欲しい本の名前や値段などの情報を管理するAPIを作成してみた！

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
│   └── views
│       └── responses.go
├── go.mod
├── go.sum
├── .env
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