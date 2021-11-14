package logUtils

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"
)

type MyLogger struct {
}

func (logger *MyLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	log.WithFields(map[string]interface{}{
		"ctx":   ctx,
		"begin": begin,
		"fc":    fc,
		"err":   err,
	}).Trace()
}

// Printf 打印数据到日志
func (logger *MyLogger) Printf(format string, v ...interface{}) {
	var logContent string
	logContent = fmt.Sprintf(format, v...)
	log.Info(logContent)
}
