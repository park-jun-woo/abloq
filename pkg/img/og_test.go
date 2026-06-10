//ff:func feature=image type=generator control=sequence
//ff:what OG가 spec.Out에 디코드 가능한 1200×630 WebP를 기록하는지 검증
package img

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/image/webp"
)

func TestOG(t *testing.T) {
	out := filepath.Join(t.TempDir(), "images", "card.webp")
	if err := OG(OGSpec{Title: "Quest", Out: out}); err != nil {
		t.Fatalf("OG: %v", err)
	}
	f, err := os.Open(out)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer f.Close()
	m, err := webp.Decode(f)
	if err != nil {
		t.Fatalf("webp decode: %v", err)
	}
	if b := m.Bounds(); b.Dx() != 1200 || b.Dy() != 630 {
		t.Errorf("bounds = %v, want 1200x630", b)
	}
	if err := OG(OGSpec{Title: "x", FontPath: "/nonexistent.ttf", Out: out}); err == nil {
		t.Error("OG with missing font expected error, got nil")
	}
}
