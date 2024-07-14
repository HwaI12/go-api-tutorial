# go-api-tutorial

簡単にGo言語とMySQLで欲しい本の名前や値段などの情報を管理するAPIを作成してみた！
## メモ
[memo.md](https://github.com/HwaI12/go-api-tutorial/blob/main/memo.md)

## 利用方法
1. リポジトリをクローン
    ```sh
    git clone https://github.com/HwaI12/go-api-tutorial.git
    ```
2. .envファイルを追加
   1. [.envファイル](#envファイル)に記載
3. テーブルを作成
   1. [MySQL](https://github.com/HwaI12/go-api-tutorial/blob/main/memo.md#mysql)に記載
4. サーバを起動
    ```sh
    go run cmd/myapp/main.go
    ```
5. curlコマンドを実行
   1. データの挿入
        ```sh
        curl -X POST http://localhost:8080/books -H \
        "Content-Type: application/json" \
                -H "X-API-KEY: <API_KEY>" \
                -d '{
            "name": "残響のテロル",
            "price": 3035
        }'
        ```
    2. 全てのデータの取得
        ```sh
        curl -X GET http://localhost:8080/books \
            -H "Content-Type: application/json" \
            -H "X-API-KEY: <API_KEY>"
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