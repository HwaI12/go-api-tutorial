package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/HwaI12/go-api-tutorial/internal/transaction"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Format(time.RFC3339)
	trnID, _ := entry.Data["trn_id"].(string)
	trnTime, _ := entry.Data["trn_time"].(string)
	logMessage := fmt.Sprintf("%s - %s - %s - %s - %s\n",
		timestamp, entry.Level.String(), trnID, trnTime, entry.Message)
	return []byte(logMessage), nil
}

func InitializeLogger() {
	Log = logrus.New()
	Log.SetFormatter(&CustomFormatter{})
	Log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/testlogfile.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
	})
}

func WithTransaction(ctx context.Context) *logrus.Entry {
	trnID, ok := ctx.Value(transaction.TrnIDKey).(string)
	if !ok {
		trnID = "unknown"
	}
	trnTime, ok := ctx.Value(transaction.TrnTimeKey).(string)
	if !ok {
		trnTime = "unknown"
	}
	return Log.WithFields(logrus.Fields{
		"trn_id":   trnID,
		"trn_time": trnTime,
	})
}
