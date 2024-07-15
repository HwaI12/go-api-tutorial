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
func InitializeGlobalTransaction() {
	globalTransaction = &TransactionInfo{
		TrnID:   uuid.New().String(),
		TrnTime: time.Now().Format(time.RFC3339),
	}
}

// 現在のグローバルトランザクション情報を取得する
func GetGlobalTransaction() *TransactionInfo {
	return globalTransaction
}

// コンテキストにトランザクション情報を設定する。
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
