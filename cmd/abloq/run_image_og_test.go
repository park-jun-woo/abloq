//ff:func feature=cli type=command control=sequence
//ff:what runImageOG가 {out}/{slug}.webp를 기록하고 front matter 참조 경로를 안내하는지 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunImageOG(t *testing.T) {
	dir := t.TempDir()
	var out bytes.Buffer
	if err := runImageOG(&out, "card", "Title Text", "Brand", "", dir); err != nil {
		t.Fatalf("runImageOG: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dir, "card.webp")); err != nil {
		t.Errorf("card.webp missing: %v", err)
	}
	if !strings.Contains(out.String(), `image: "/images/card.webp"`) {
		t.Errorf("want front matter hint, got %q", out.String())
	}
	if err := runImageOG(&out, "card", "T", "", "/nonexistent.ttf", dir); err == nil {
		t.Error("missing font must error")
	}
}
