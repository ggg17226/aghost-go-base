package configUtils

import (
	"fmt"
	"github.com/ggg17226/aghost-go-base/pkg/utils/fileUtils"
	"github.com/ggg17226/aghost-go-base/pkg/utils/logUtils"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

var (
	configFileName = "app"
	configFilePath = make([]string, 0)
	configFileType = [...]string{"yaml", "yml", "json", "toml"}
	ConfigKeyList  = make([][]string, 0)
)

// BindEnvConfigKey 传入要从环境变量中读取的key列表，否则不会读取环境变量中的配置
func BindEnvConfigKey(keyList *[][]string) {
	if len(*keyList) < 1 {
		return
	}
	for _, s := range *keyList {
		err := viper.BindEnv(s...)
		if err != nil {
			panic("viper bind env error")
		}
	}
}

// SetConfigFileName 设置配置文件的文件名
func SetConfigFileName(fileName string) {
	configFileName = fileName
}

// AppendConfigFilePath 添加读取配置文件的路径
func AppendConfigFilePath(pathList ...string) {
	if len(pathList) < 1 {
		return
	}
	for _, s := range pathList {
		configFilePath = append(configFilePath, s)
	}
}

// 添加日志用配置key
func initLogConfigKeyList() {
	ConfigKeyList = append(ConfigKeyList,
		[]string{logUtils.LogPathKey, logUtils.LogPathEnvKey},
		[]string{logUtils.LogTargetKey, logUtils.LogTargetEnvKey},
		[]string{logUtils.LogLevelKey, logUtils.LogLevelEnvKey},
		[]string{logUtils.LogFilenameKey, logUtils.LogFilenameEnvKey},
		[]string{logUtils.LogFormatKey, logUtils.LogFormatEnvKey},
		[]string{logUtils.LogMaxAgeKey, logUtils.LogMaxAgeEnvKey},
		[]string{logUtils.LogMaxCountKey, logUtils.LogMaxCountEnvKey},
	)
}

// 添加默认的配置文件路径
func initConfigPath() {
	configFilePath = append(configFilePath, "/etc/")
	configFilePath = append(configFilePath, "./")
	configFilePath = append(configFilePath, "./conf/")
	configFilePath = append(configFilePath, "/etc/"+configFileName+"/")
}

// InitConfigAndLog 读取配置
// 读取顺序： 环境变量 -> 自己添加的路径 -> /etc/${configFileName}/ -> /etc -> ./
// yaml->yml->json->toml
// 以先读到的为准
func InitConfigAndLog() {
	initLogConfigKeyList()
	BindEnvConfigKey(&ConfigKeyList)
	viper.AllowEmptyEnv(false)
	viper.AutomaticEnv()

	initConfigPath()

	for _, path := range configFilePath {
		for _, typeName := range configFileType {
			absFilePath, err := filepath.Abs(path + configFileName + "." + typeName)
			if err != nil {
				continue
			}
			if fileUtils.FileExists(absFilePath) {
				if typeName == "yml" {
					viper.SetConfigType("yaml")
				} else {
					viper.SetConfigType(typeName)
				}
				f, err := os.Open(absFilePath)
				if err != nil {
					continue
				}
				if err := viper.MergeConfig(f); err != nil {
					if _, ok := err.(viper.ConfigFileNotFoundError); ok {
					} else {
						panic(fmt.Errorf("Fatal error reding config file: %w \n", err))
					}
				}
			}
		}
	}
	initLog()
}

// 初始化日志
func initLog() {
	viper.SetDefault(logUtils.LogPathKey, logUtils.DefaultLogDir)
	viper.SetDefault(logUtils.LogTargetKey, logUtils.LogDefaultTarget)
	viper.SetDefault(logUtils.LogLevelKey, logUtils.LogDefaultOutputLevelName)
	viper.SetDefault(logUtils.LogFilenameKey, logUtils.DefaultLogFilename)
	viper.SetDefault(logUtils.LogFormatKey, logUtils.DefaultLogFormat)

	logUtils.LogPath = viper.GetString(logUtils.LogPathKey)
	logUtils.LogTarget = viper.GetString(logUtils.LogTargetKey)
	logUtils.LogLevel = viper.GetString(logUtils.LogLevelKey)
	logUtils.LogFilename = viper.GetString(logUtils.LogFilenameKey)
	logUtils.LogFormat = viper.GetString(logUtils.LogFormatKey)

	logUtils.LogMaxAge = viper.GetDuration(logUtils.LogMaxAgeKey)
	logUtils.LogMaxCount = viper.GetInt(logUtils.LogMaxCountKey)

	logUtils.InitLog()
}
