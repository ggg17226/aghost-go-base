package fileUtils

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// FileExists 检查文件或文件夹是否存在(check file or dir is exists)
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else {
			return false
		}
	}
	return true
}

// GenerateFileNameWithTimeAndRandomNumber 根据时间和随机数生成文件名(generate file name with time and random number)
func GenerateFileNameWithTimeAndRandomNumber() string {
	timeStr := time.Now().Format("2006-01-02T15-04-05Z0700")
	return timeStr + "-" + strconv.Itoa(rand.Intn(89999)+10000)
}

// CreateAndWriteFile 创建并写入文件(create file and write content)
func CreateAndWriteFile(path string, content []byte) error {
	absPath, absPathErr := filepath.Abs(path)
	if absPathErr != nil {
		return absPathErr
	}
	parentPath := filepath.Dir(absPath)
	makeParentDirErr := os.MkdirAll(parentPath, 755)
	if makeParentDirErr != nil {
		return makeParentDirErr
	}
	fp, openFileErr := os.Create(absPath)
	if openFileErr != nil {
		return openFileErr
	}
	defer func() {
		err := fp.Close()
		if err != nil {
			log.WithField("path", path).
				WithField("err", err).
				Error("close file error")
		}
	}()
	_, writeErr := fp.Write(content)
	if writeErr != nil {
		return writeErr
	}
	return nil
}
