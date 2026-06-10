//ff:func feature=gate type=parser control=sequence topic=baseline
//ff:what 테스트 저장소에 상대 경로 파일을 기록 (디렉토리 자동 생성)
package gate

import (
	"os"
	"path/filepath"
	"testing"
)

func writeRepoFile(t *testing.T, dir, rel, content string) {
	t.Helper()
	path := filepath.Join(dir, rel)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
