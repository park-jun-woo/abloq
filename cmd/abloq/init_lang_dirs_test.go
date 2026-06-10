//ff:func feature=init type=command control=iteration dimension=1
//ff:what initLangDirs가 언어 1개의 전 섹션 디렉토리를 .gitkeep과 함께 만드는지 검증
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitLangDirs(t *testing.T) {
	dir := t.TempDir()
	if err := initLangDirs(dir, "ko", []string{"opinion", "tech"}); err != nil {
		t.Fatalf("initLangDirs: %v", err)
	}
	for _, section := range []string{"opinion", "tech"} {
		p := filepath.Join(dir, "content", "ko", section, ".gitkeep")
		if _, err := os.Stat(p); err != nil {
			t.Errorf("missing %s: %v", p, err)
		}
	}
	blocked := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(blocked, []byte(""), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := initLangDirs(blocked, "ko", []string{"tech"}); err == nil {
		t.Error("initLangDirs under a regular file expected error, got nil")
	}
}
