//ff:func feature=init type=generator control=sequence
//ff:what fs.FS의 파일 1개를 대상 경로에 기록 — 상위 디렉토리 생성 포함, 같은 바이트 재기록이라 멱등
package scaffold

import (
	"io/fs"
	"os"
	"path/filepath"
)

// CopyFile writes one embedded file to dstPath, creating parent directories.
func CopyFile(fsys fs.FS, src, dstPath string) error {
	data, err := fs.ReadFile(fsys, src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dstPath, data, 0o644)
}
