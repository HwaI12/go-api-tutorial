package Log

import (
	"fmt"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 現在のタイムスタンプをRFC3339形式で取得
	timestamp := time.Now().Format(time.RFC3339)

	// ログエントリからトランザクションIDとトランザクション時間を取得
	trnID, _ := entry.Data["trn_id"].(string)
	trnTime, _ := entry.Data["trn_time"].(string)

	// ログメッセージをカスタムフォーマットで生成
	logMessage := fmt.Sprintf("%s - %s - %s - %s - %s\n",
		timestamp, entry.Level.String(), trnID, trnTime, entry.Message)

	// フォーマットされたログメッセージをバイトスライスとして返す
	return []byte(logMessage), nil
}

func InitializeLogger() {
	// ログのフォーマットと出力先を設定
	logrus.SetFormatter(&CustomFormatter{})
	// ログの出力先を設定
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "logs/testlogfile.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
	})
}
