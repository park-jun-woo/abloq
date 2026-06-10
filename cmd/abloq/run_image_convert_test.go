//ff:func feature=cli type=command control=sequence
//ff:what runImageConvert가 slug 기본값(원본 파일명)으로 WebP를 만들고 마크다운 참조를 안내하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunImageConvert(t *testing.T) {
	dir := t.TempDir()
	src := writePNGFixture(t, dir)
	var out bytes.Buffer
	if err := runImageConvert(&out, src, "", dir, 1400); err != nil {
		t.Fatalf("runImageConvert: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "in.webp")); err != nil {
		t.Errorf("in.webp missing (slug default = filename): %v", err)
	}
	if !strings.Contains(out.String(), "![alt text](/images/in.webp)") {
		t.Errorf("want markdown hint, got %q", out.String())
	}
	if err := runImageConvert(&out, filepath.Join(dir, "missing.png"), "", dir, 1400); err == nil {
		t.Error("missing source must error")
	}
}
