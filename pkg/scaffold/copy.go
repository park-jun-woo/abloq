//ff:func feature=init type=generator control=iteration dimension=1
//ff:what 템플릿 fs.FS 전체를 대상 디렉토리에 복제(degit식) — 복사한 파일 수 반환
package scaffold

import (
	"io/fs"
	"path/filepath"
)

// Copy clones every file of fsys into dst, preserving relative paths.
func Copy(fsys fs.FS, dst string) (int, error) {
	paths, err := ListDir(fsys, ".")
	if err != nil {
		return 0, err
	}
	for i, p := range paths {
		if err := CopyFile(fsys, p, filepath.Join(dst, p)); err != nil {
			return i, err
		}
	}
	return len(paths), nil
}
