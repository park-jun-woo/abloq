//ff:type feature=init type=parser
//ff:what 테스트 픽스처 — 루트 외 디렉토리의 ReadDir을 실패시키는 fs.FS, ListDir 재귀 에러 분기 검증용
package scaffold

import (
	"io/fs"
	"testing/fstest"
)

// failSubFS fails ReadDir for every non-root directory.
type failSubFS struct{ fstest.MapFS }

func (f failSubFS) ReadDir(name string) ([]fs.DirEntry, error) {
	if name != "." {
		return nil, fs.ErrPermission
	}
	return f.MapFS.ReadDir(name)
}
