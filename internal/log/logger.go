package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/HwaI12/go-api-tutorial/internal/transaction"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

// カスタムログフォーマッタを定義する構造体
type CustomFormatter struct{}

// ログエントリをカスタムフォーマットでフォーマットする。
// エントリのタイムスタンプ、レベル、トランザクションID、トランザクション時間、メッセージを含む。
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format(time.RFC3339)
	trnID, _ := entry.Data["trn_id"].(string)
	trnTime, _ := entry.Data["trn_time"].(string)
	logMessage := fmt.Sprintf("%s - %s - %s - %s - %s\n",
		timestamp, entry.Level.String(), trnID, trnTime, entry.Message)
	return []byte(logMessage), nil
}

// ログの初期設定を行う。
// カスタムフォーマッタを設定し、ログの出力先をファイルにする。
// ファイルのローテーションも設定する。
func InitializeLogger() {
	logrus.SetFormatter(&CustomFormatter{})
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "logs/testlogfile.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
	})
	logrus.SetLevel(logrus.DebugLevel) // Debugレベルのログも記録するよう設定
}

// トランザクション情報を含むログエントリを作成する。
// コンテキストからトランザクションIDとトランザクション時間を取得し、フィールドとして追加する。
func WithTransaction(ctx context.Context) *logrus.Entry {
	trnID, _ := ctx.Value(transaction.TrnIDKey).(string)
	trnTime, _ := ctx.Value(transaction.TrnTimeKey).(string)
	return logrus.WithFields(logrus.Fields{
		"trn_id":   trnID,
		"trn_time": trnTime,
	})
}
