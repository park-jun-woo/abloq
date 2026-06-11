//ff:func feature=cli type=command control=sequence
//ff:what runImageOGLocal 검증 — {out}/{slug}.webp 결정론 렌더 기록과 front matter 참조 안내, 렌더 실패 전파
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunImageOGLocal(t *testing.T) {
	dir := t.TempDir()
	var out bytes.Buffer
	opts := imageOGOpts{Slug: "card", Title: "Title Text", Brand: "Brand", OutDir: dir}
	if err := runImageOGLocal(&out, opts); err != nil {
		t.Fatalf("runImageOGLocal: %v", err)
	}
	dst := filepath.Join(dir, "card.webp")
	if _, err := os.Stat(dst); err != nil {
		t.Fatalf("card.webp missing: %v", err)
	}
	if !strings.Contains(out.String(), dst) ||
		!strings.Contains(out.String(), `front matter: image: "/images/card.webp"`) {
		t.Errorf("output = %q, want path + front matter hint", out.String())
	}

	// render failure propagates
	bad := imageOGOpts{Slug: "card", Title: "T", FontPath: "/nonexistent.ttf", OutDir: dir}
	if err := runImageOGLocal(&out, bad); err == nil {
		t.Error("missing font must error")
	}
}
