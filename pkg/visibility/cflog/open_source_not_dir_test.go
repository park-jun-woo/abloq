//ff:func feature=visibility type=client control=sequence topic=crawl
//ff:what OpenSource가 일반 파일을 소스로 거부하는지 검증
package cflog

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenSourceNotDir(t *testing.T) {
	f := filepath.Join(t.TempDir(), "file.txt")
	if err := os.WriteFile(f, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenSource(f); err == nil {
		t.Error("plain file accepted as source")
	}
}
