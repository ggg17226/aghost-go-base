package test

import (
	"github.com/ggg17226/aghost-go-base/pkg/utils/configUtils"
	"github.com/ggg17226/aghost-go-base/pkg/utils/randomUtils"
	"github.com/spf13/viper"
	"testing"
)

const (
	configPathKey       = "path"
	envPathKey          = "PATH"
	testRandomStrLength = 32
)

func TestConfigAndLog(t *testing.T) {
	configUtils.ConfigKeyList = append(configUtils.ConfigKeyList, []string{configPathKey, envPathKey})
	configUtils.InitConfigAndLog()
	if len(viper.GetString(configPathKey)) < 1 {
		t.Fatalf("test failed with get path from env")
	}
}

func TestRandomStr(t *testing.T) {
	randString1 := randomUtils.RandString()
	if len(randString1) != randomUtils.DefaultLength {
		t.Fatalf("test failed with get default length random string")
	}
	randString2 := randomUtils.RandString()
	if randString1 == randString2 {
		t.Fatalf("test failed with get same random string")
	}
	if len(randomUtils.RandStringWithLength(testRandomStrLength)) != testRandomStrLength {
		t.Fatalf("test failed with get random string")
	}

	randString3 := randomUtils.RandReadableString()
	if len(randString3) != randomUtils.DefaultLength {
		t.Fatalf("test failed with get default length random string")
	}
	randString4 := randomUtils.RandReadableString()
	if randString3 == randString4 {
		t.Fatalf("test failed with get same random string")
	}
	if len(randomUtils.RandReadableStringWithLength(testRandomStrLength)) != testRandomStrLength {
		t.Fatalf("test failed with get random string")
	}

}
