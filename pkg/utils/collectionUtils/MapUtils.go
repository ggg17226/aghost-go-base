package collectionUtils

import (
	"net/http"
	"sort"
	"strings"
)

// GetKeyListFromHeaderMap 从http头map中获取有序的key列表
func GetKeyListFromHeaderMap(m *http.Header) []string {
	keys := make([]string, 0, len(*m))
	for k := range *m {
		keys = append(keys, strings.ToLower(k))
	}
	sort.Strings(keys)
	return keys
}
