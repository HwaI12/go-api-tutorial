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

func InitializeTransaction(ctx context.Context) context.Context {
	trnID := uuid.New().String()
	trnTime := time.Now().Format(time.RFC3339)

	// コンテキストにトランザクションIDとトランザクション時間を設定
	ctx = context.WithValue(ctx, TrnIDKey, trnID)
	ctx = context.WithValue(ctx, TrnTimeKey, trnTime)
	return ctx
}
