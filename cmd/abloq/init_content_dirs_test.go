//ff:func feature=init type=command control=iteration dimension=1
//ff:what initContentDirs가 언어×섹션 콘텐츠 디렉토리와 quests/queue를 .gitkeep과 함께 만드는지 검증
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitContentDirs(t *testing.T) {
	dir := t.TempDir()
	if err := initContentDirs(dir, []string{"ko", "en"}, []string{"tech"}); err != nil {
		t.Fatalf("initContentDirs: %v", err)
	}
	musts := []string{
		"content/ko/tech/.gitkeep",
		"content/en/tech/.gitkeep",
		"quests/queue/.gitkeep",
	}
	for _, m := range musts {
		if _, err := os.Stat(filepath.Join(dir, m)); err != nil {
			t.Errorf("missing %s: %v", m, err)
		}
	}
	blocked := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(blocked, []byte(""), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := initContentDirs(blocked, []string{"ko"}, []string{"tech"}); err == nil {
		t.Error("initContentDirs under a regular file expected error, got nil")
	}
}
