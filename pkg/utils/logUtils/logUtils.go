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
	LogMaxAge   = time.Duration(int64(0))
	LogMaxCount = int(0)
)

// InitLog 初始化日志
func InitLog() {

	if LogTarget != LogTargetStd && !fileUtils.FileExists(LogPath) {
		err := os.Mkdir(LogPath, 0755)
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

	rotateLogOptions := make([]rotateLogs.Option, 0)

	if runtime.GOOS != "windows" {
		rotateLogOptions = append(rotateLogOptions, rotateLogs.WithLinkName(LogFilename))
	}

	rotateLogOptions = append(rotateLogOptions, rotateLogs.WithRotationTime(12*time.Hour))
	rotateLogOptions = append(rotateLogOptions, rotateLogs.WithClock(rotateLogs.Local))

	if LogMaxCount > 0 {
		rotateLogOptions = append(rotateLogOptions, rotateLogs.WithRotationCount(LogMaxCount))
	}

	if int64(LogMaxAge) > 0 {
		rotateLogOptions = append(rotateLogOptions, rotateLogs.WithMaxAge(LogMaxAge))
	}

	rotateWriter, rotateLogErr := rotateLogs.New(logFilePath+".%Y-%m-%d-%H", rotateLogOptions...)
	if rotateLogErr != nil {
		panic(rotateLogErr.Error())
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
	case LogLevelTraceName:
		log.SetLevel(LogLevelTrace)
		break
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
