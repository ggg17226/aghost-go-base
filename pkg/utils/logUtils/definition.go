package logUtils

import log "github.com/sirupsen/logrus"

const (
	// 定义日志配置使用的环境变量

	LogPathEnvKey     = "app_log_path"
	LogPathKey        = "app.log.path"
	LogTargetEnvKey   = "app_log_target"
	LogTargetKey      = "app.log.target"
	LogLevelEnvKey    = "app_log_level"
	LogLevelKey       = "app.log.level"
	LogFilenameEnvKey = "app_log_filename"
	LogFilenameKey    = "app.log.filename"
	LogFormatEnvKey   = "app_log_format"
	LogFormatKey      = "app.log.format"
	LogMaxAgeEnvKey   = "app_log_max_age"
	LogMaxAgeKey      = "app.log.max.age"
	LogMaxCountEnvKey = "app_log_max_count"
	LogMaxCountKey    = "app.log.max.count"

	// 定义日志输出格式

	LogFormatJson      = "json"
	LogFormatPlainText = "text"

	DefaultLogFormat = LogFormatJson

	// 定义日志级别常量

	LogLevelTrace     = log.TraceLevel
	LogLevelTraceName = "trace"
	LogLevelInfo      = log.InfoLevel
	LogLevelInfoName  = "info"
	LogLevelError     = log.ErrorLevel
	LogLevelErrorName = "error"
	LogLevelDebug     = log.DebugLevel
	LogLevelDebugName = "debug"
	LogLevelWarn      = log.WarnLevel
	LogLevelWarnName  = "warn"

	LogDefaultOutputLevelName = LogLevelInfoName

	// 定义日志输出目标常量

	LogTargetFile = "file"
	LogTargetStd  = "std"
	LogTargetBoth = "both"

	LogDefaultTarget = LogTargetFile

	// 定义默认日志输出文件配置

	DefaultLogDir      = "logs"
	DefaultLogFilename = "app.log"
)
