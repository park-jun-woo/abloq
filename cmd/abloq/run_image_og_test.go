//ff:func feature=cli type=command control=sequence
//ff:what runImageOG 기본(local) 경로가 {out}/{slug}.webp를 기록하고 front matter 참조를 안내하는지 — img.OG 직행과 바이트 동일 검증
package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/img"
)

func TestRunImageOG(t *testing.T) {
	dir := t.TempDir()
	var out bytes.Buffer
	opts := imageOGOpts{Slug: "card", Title: "Title Text", Brand: "Brand", OutDir: dir, Count: 1}
	if err := runImageOG(&out, opts); err != nil {
		t.Fatalf("runImageOG: %v", err)
	}
	got, err := os.ReadFile(filepath.Join(dir, "card.webp"))
	if err != nil {
		t.Fatalf("card.webp missing: %v", err)
	}
	if !strings.Contains(out.String(), `image: "/images/card.webp"`) {
		t.Errorf("want front matter hint, got %q", out.String())
	}
	// byte identity with the direct deterministic path (no Provider detour)
	ref := filepath.Join(dir, "ref.webp")
	if err := img.OG(img.OGSpec{Title: "Title Text", Brand: "Brand", Out: ref}); err != nil {
		t.Fatalf("img.OG: %v", err)
	}
	want, _ := os.ReadFile(ref)
	if !bytes.Equal(got, want) {
		t.Errorf("local path output differs from direct RenderOG bytes (%d vs %d)", len(got), len(want))
	}
	bad := imageOGOpts{Slug: "card", Title: "T", FontPath: "/nonexistent.ttf", OutDir: dir, Count: 1}
	if err := runImageOG(&out, bad); err == nil {
		t.Error("missing font must error")
	}
}
