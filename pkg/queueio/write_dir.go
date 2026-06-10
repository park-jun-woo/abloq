//ff:func feature=queueio type=generator control=iteration dimension=1
//ff:what 큐 디렉토리에 항목 파일 일괄 기록 — 결정적 직렬화라 같은 입력이면 바이트 동일(멱등), CLI와 exporter가 공유
package queueio

import (
	"os"
	"path/filepath"
)

// WriteDir writes one file per item under dir (created if missing). It never
// deletes files: deletion is the agent's consumption signal, which the next
// export cycle reads back as status → consumed.
func WriteDir(dir string, items []Item) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for _, it := range items {
		path := filepath.Join(dir, Filename(it))
		if err := os.WriteFile(path, Serialize(it), 0o644); err != nil {
			return err
		}
	}
	return nil
}
