//ff:func feature=queueio type=generator control=sequence
//ff:what WriteDir가 디렉토리를 만들고 항목 파일을 기록하며 기존 파일을 지우지 않는지(에이전트 삭제 신호 보존) 검증
package queueio

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "quests", "queue")
	keep := []byte("agent-owned\n")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "other.yaml"), keep, 0o644); err != nil {
		t.Fatal(err)
	}
	it := Item{Kind: "refresh", Slug: "post-a", Lang: "ko", Section: "tech"}
	if err := WriteDir(dir, []Item{it}); err != nil {
		t.Fatalf("WriteDir: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, Filename(it)))
	if err != nil {
		t.Fatalf("queue file missing: %v", err)
	}
	if string(data) != string(Serialize(it)) {
		t.Error("file content must equal Serialize output")
	}
	if got, _ := os.ReadFile(filepath.Join(dir, "other.yaml")); string(got) != string(keep) {
		t.Error("WriteDir must never touch unrelated files")
	}
	// A file standing where the directory should be surfaces as an error.
	blocked := filepath.Join(t.TempDir(), "not-a-dir")
	if err := os.WriteFile(blocked, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := WriteDir(blocked, nil); err == nil {
		t.Error("MkdirAll over a file must error")
	}
	if err := WriteDir(filepath.Join(blocked, "sub"), []Item{it}); err == nil {
		t.Error("unwritable destination must error")
	}
}
