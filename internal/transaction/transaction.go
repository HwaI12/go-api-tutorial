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

var globalTransaction *TransactionInfo

func InitializeGlobalTransaction() {
	globalTransaction = &TransactionInfo{
		TrnID:   uuid.New().String(),
		TrnTime: time.Now().Format(time.RFC3339),
	}
}

func GetGlobalTransaction() *TransactionInfo {
	return globalTransaction
}

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
