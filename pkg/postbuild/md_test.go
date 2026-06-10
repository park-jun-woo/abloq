//ff:func feature=postbuild type=generator control=sequence
//ff:what MD가 글마다 public/ 옆 .md를 기록하고 개수를 반환하는지, 재실행이 멱등인지 검증
package postbuild

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestMD(t *testing.T) {
	b := &blogyaml.Blog{Languages: []string{"ko"}}
	b.Site.DefaultLangInSubdir = true
	dir := t.TempDir()
	post := filepath.Join(dir, "content", "ko", "tech", "hello.md")
	if err := os.MkdirAll(filepath.Dir(post), 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	src := "---\ntitle: \"안녕\"\ntags: [a]\n---\n\n본문.\n"
	if err := os.WriteFile(post, []byte(src), 0o644); err != nil {
		t.Fatalf("write post: %v", err)
	}
	n, err := MD(dir, b)
	if err != nil || n != 1 {
		t.Fatalf("MD = %d, %v; want 1, nil", n, err)
	}
	got, err := os.ReadFile(filepath.Join(dir, "public", "ko", "tech", "hello.md"))
	if err != nil || string(got) != "# 안녕\n\n본문.\n" {
		t.Errorf("served md = %q, err %v", got, err)
	}
	if n, err := MD(dir, b); err != nil || n != 1 {
		t.Errorf("second MD = %d, %v; want 1, nil", n, err)
	}
	blocked := t.TempDir()
	post2 := filepath.Join(blocked, "content", "ko", "tech", "x.md")
	if err := os.MkdirAll(filepath.Dir(post2), 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(post2, []byte("x\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	if err := os.WriteFile(filepath.Join(blocked, "public"), []byte(""), 0o644); err != nil {
		t.Fatalf("write public file: %v", err)
	}
	if _, err := MD(blocked, b); err == nil {
		t.Error("MD with public/ blocked by a regular file expected error, got nil")
	}
}
