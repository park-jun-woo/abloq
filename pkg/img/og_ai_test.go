//ff:func feature=image type=generator control=iteration dimension=1
//ff:what OGAI가 단일 직행/다중 안 드래프트 경로·안×count 전개·오버레이·부분 실패 집계를 수행하는지 stub Provider로 검증
package img

import (
	"context"
	"errors"
	"image/color"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/image/webp"
)

func TestOGAI(t *testing.T) {
	dir := t.TempDir()
	blue := stubOGProvider{w: 900, h: 900, c: color.NRGBA{10, 10, 200, 255}}

	// single direct call: straight to OutDir/{slug}.webp, 1200x630.
	spec := OGAISpec{Slug: "post", OutDir: filepath.Join(dir, "static"), DraftDir: filepath.Join(dir, "og"), Count: 1}
	outcomes := OGAI(context.Background(), spec, []OGVariant{{Name: "default", Model: "m1", Prompt: "p", Provider: blue}})
	if len(outcomes) != 1 || outcomes[0].Err != nil {
		t.Fatalf("direct outcomes = %+v", outcomes)
	}
	f, err := os.Open(filepath.Join(dir, "static", "post.webp"))
	if err != nil {
		t.Fatalf("direct output missing: %v", err)
	}
	m, err := webp.Decode(f)
	f.Close()
	if err != nil {
		t.Fatalf("webp decode: %v", err)
	}
	if b := m.Bounds(); b.Dx() != 1200 || b.Dy() != 630 {
		t.Errorf("bounds = %v, want 1200x630", b)
	}

	// multi path: 2 variants x count 2 = 4 drafts named {variant}-{n}.webp.
	multi := OGAISpec{Slug: "post", Title: "T", OutDir: filepath.Join(dir, "static"), DraftDir: filepath.Join(dir, "og"), Multi: true, Count: 2}
	variants := []OGVariant{
		{Name: "minimal", Model: "m1", Prompt: "p1", Provider: blue},
		{Name: "photo", Model: "m2", Prompt: "p2", Overlay: true, Provider: blue},
	}
	outcomes = OGAI(context.Background(), multi, variants)
	if len(outcomes) != 4 {
		t.Fatalf("multi outcomes = %d, want 4", len(outcomes))
	}
	for _, name := range []string{"minimal-1.webp", "minimal-2.webp", "photo-1.webp", "photo-2.webp"} {
		if _, err := os.Stat(filepath.Join(dir, "og", "post", name)); err != nil {
			t.Errorf("draft %s missing: %v", name, err)
		}
	}
	plain, _ := os.ReadFile(filepath.Join(dir, "og", "post", "minimal-1.webp"))
	overlaid, _ := os.ReadFile(filepath.Join(dir, "og", "post", "photo-1.webp"))
	if string(plain) == string(overlaid) {
		t.Error("overlay variant equals plain variant — title was not composited")
	}

	// partial failure: failing variant keeps the successful sibling on disk.
	boom := stubOGProvider{err: errors.New("quota")}
	fail := OGAISpec{Slug: "mix", OutDir: dir, DraftDir: filepath.Join(dir, "og"), Multi: true, Count: 1}
	outcomes = OGAI(context.Background(), fail, []OGVariant{
		{Name: "ok", Model: "m1", Prompt: "p", Provider: blue},
		{Name: "bad", Model: "m1", Prompt: "p", Provider: boom},
	})
	if outcomes[0].Err != nil || outcomes[1].Err == nil {
		t.Fatalf("partial outcomes = %+v", outcomes)
	}
	if _, err := os.Stat(filepath.Join(dir, "og", "mix", "ok-1.webp")); err != nil {
		t.Errorf("successful sibling lost: %v", err)
	}
}
