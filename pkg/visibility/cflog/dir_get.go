//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what 디렉토리 소스에서 키 1개의 파일 스트림 열기
package cflog

import (
	"io"
	"os"
	"path/filepath"
)

// Get opens one log file by key.
func (s DirSource) Get(key string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(s.Root, key))
}
