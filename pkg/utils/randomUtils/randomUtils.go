package randomUtils

import (
	"math/rand"
	"time"
)

var (
	source = rand.NewSource(time.Now().UnixNano())
)

// RandStringWithLength 生成随机字符串
func RandStringWithLength(n int) string {
	return randomWithLength(n, letterBytes)
}

// RandString 成随机16个字符的字符串
func RandString() string {
	return RandStringWithLength(DefaultLength)
}

func RandReadableStringWithLength(n int) string {
	return randomWithLength(n, readableLetterBytes)
}

func RandReadableString() string {
	return RandReadableStringWithLength(DefaultLength)
}

func randomWithLength(n int, characterSet string) string {
	b := make([]byte, n)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, source.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = source.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(characterSet) {
			b[i] = characterSet[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
