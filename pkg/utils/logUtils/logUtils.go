package logUtils

import (
	"github.com/ggg17226/aghost-go-base/pkg/utils/fileUtils"
	rotateLogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	LogPath     = DefaultLogDir
	LogTarget   = LogDefaultTarget
	LogLevel    = LogDefaultOutputLevelName
	LogFilename = DefaultLogFilename
	LogFormat   = DefaultLogFormat
)

// InitLog 初始化日志
func InitLog() {

	if LogTarget != LogTargetStd && !fileUtils.FileExists(LogPath) {
		err := os.Mkdir(LogPath, 0777)
		if err != nil {
			panic("create log dir error")
		}
	}

	switch LogFormat {
	default:
	case LogFormatJson:
		// 设置日志格式为json格式
		log.SetFormatter(&log.JSONFormatter{PrettyPrint: false})
		break
	case LogFormatPlainText:
		// 设置日志格式为文本格式
		log.SetFormatter(&log.TextFormatter{})
		break
	}

	/**
	设置日志格式
	*/
	logFilePath := filepath.Join(LogPath, LogFilename)

	var rotateWriter *rotateLogs.RotateLogs
	if runtime.GOOS == "windows" {
		rotateWriter, _ = rotateLogs.New(
			logFilePath+".%Y-%m-%d-%H",
			rotateLogs.WithRotationTime(12*time.Hour),
			rotateLogs.WithClock(rotateLogs.Local),
		)
	} else {
		rotateWriter, _ = rotateLogs.New(
			logFilePath+".%Y-%m-%d-%H",
			rotateLogs.WithLinkName(LogFilename),
			rotateLogs.WithRotationTime(12*time.Hour),
			rotateLogs.WithClock(rotateLogs.Local),
		)
	}

	/**
	  配置日志输出目标
	*/
	switch LogTarget {
	default:
	case LogTargetStd:
		log.SetOutput(os.Stdout)
		break
	case LogTargetBoth:
		mw := io.MultiWriter(os.Stdout, rotateWriter)
		log.SetOutput(mw)
		break
	case LogTargetFile:
		log.SetOutput(rotateWriter)
		break
	}

	switch LogLevel {
	default:
	case LogLevelInfoName:
		log.SetLevel(LogLevelInfo)
		break
	case LogLevelErrorName:
		log.SetLevel(LogLevelError)
		break
	case LogLevelDebugName:
		log.SetLevel(LogLevelDebug)
		break
	case LogLevelWarnName:
		log.SetLevel(LogLevelWarn)
		break
	}

}
