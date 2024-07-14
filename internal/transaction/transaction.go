package transaction

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ctxKey string

const (
	TrnIDKey   ctxKey = "trn_id"
	TrnTimeKey ctxKey = "trn_time"
)

type TransactionInfo struct {
	TrnID   string
	TrnTime string
}

// グローバルトランザクション情報を保持する変数
var globalTransaction *TransactionInfo

// グローバルトランザクション情報を初期化する
// トランザクションIDとトランザクション時間を生成して設定する
func InitializeGlobalTransaction() {
	globalTransaction = &TransactionInfo{
		TrnID:   uuid.New().String(),
		TrnTime: time.Now().Format(time.RFC3339),
	}
}

// 現在のグローバルトランザクション情報を取得する
// グローバルトランザクション情報のポインタを返す
func GetGlobalTransaction() *TransactionInfo {
	return globalTransaction
}

// コンテキストにトランザクション情報を設定する。
// グローバルトランザクション情報が存在しない場合、初期化
// トランザクションIDとトランザクション時間をコンテキストに設定して返す
func InitializeTransaction(ctx context.Context) context.Context {
	if globalTransaction == nil {
		InitializeGlobalTransaction()
	}
	trnID := globalTransaction.TrnID
	trnTime := globalTransaction.TrnTime
	ctx = context.WithValue(ctx, TrnIDKey, trnID)
	ctx = context.WithValue(ctx, TrnTimeKey, trnTime)
	return ctx
}
