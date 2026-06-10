//ff:func feature=visibility type=client control=iteration dimension=1 topic=crawl
//ff:what 디렉토리의 일반 파일명을 정렬된 키 목록으로 — prefix 필터와 afterKey 초과분만
package cflog

import (
	"os"
	"sort"
	"strings"
)

// List returns the directory's regular file names as keys, ascending,
// filtered by prefix and strictly after afterKey.
func (s DirSource) List(prefix, afterKey string) ([]string, error) {
	entries, err := os.ReadDir(s.Root)
	if err != nil {
		return nil, err
	}
	var keys []string
	for _, e := range entries {
		name := e.Name()
		if !e.IsDir() && strings.HasPrefix(name, prefix) && name > afterKey {
			keys = append(keys, name)
		}
	}
	sort.Strings(keys)
	return keys, nil
}
