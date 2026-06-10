//ff:func feature=cli type=command control=sequence topic=diagnostics
//ff:what loadValidBlog가 유효 blog.yaml을 반환하고 누락·검증 실패 시 진단 출력과 함께 에러를 내는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadValidBlog(t *testing.T) {
	var out bytes.Buffer
	b, err := loadValidBlog(&out, writeBlogFixture(t))
	if err != nil || b == nil {
		t.Fatalf("loadValidBlog on fixture: %v", err)
	}
	if _, err := loadValidBlog(&out, t.TempDir()); err == nil {
		t.Errorf("want IO error for missing blog.yaml, got nil")
	}
	invalid := t.TempDir()
	if err := os.WriteFile(filepath.Join(invalid, "blog.yaml"), []byte("languages: [ko]\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	out.Reset()
	if _, err := loadValidBlog(&out, invalid); err == nil {
		t.Fatalf("want validation error, got nil")
	}
	if !strings.Contains(out.String(), "[") {
		t.Errorf("want diagnostics printed, got %q", out.String())
	}
}
