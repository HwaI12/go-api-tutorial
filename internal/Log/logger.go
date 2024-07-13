package logger

import (
	"fmt"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

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
	logrus.SetFormatter(&CustomFormatter{})
	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "logs/testlogfile.log",
		MaxSize:    500,
		MaxBackups: 3,
		MaxAge:     28,
	})
}
