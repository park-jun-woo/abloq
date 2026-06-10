//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what 파생물 목록을 블로그 루트 아래에 기록 — 상위 디렉토리 생성 포함, 같은 바이트 재기록이라 멱등
package gen

import (
	"os"
	"path/filepath"
)

// Write persists every derived file under dir, creating parent directories.
func Write(dir string, outs []Output) error {
	for _, o := range outs {
		path := filepath.Join(dir, o.Path)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(path, o.Data, 0o644); err != nil {
			return err
		}
	}
	return nil
}
