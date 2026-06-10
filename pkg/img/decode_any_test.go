//ff:func feature=image type=parser control=sequence
//ff:what DecodeAny가 PNG를 디코드하고 없는 파일에 에러를 내는지 검증
package img

import (
	"image/color"
	"path/filepath"
	"testing"
)

func TestDecodeAny(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "in.png")
	writeTestPNG(t, src, 8, 4, color.NRGBA{255, 0, 0, 255})
	m, err := DecodeAny(src)
	if err != nil {
		t.Fatalf("DecodeAny: %v", err)
	}
	if b := m.Bounds(); b.Dx() != 8 || b.Dy() != 4 {
		t.Errorf("bounds = %v, want 8x4", b)
	}
	if _, err := DecodeAny(filepath.Join(dir, "missing.png")); err == nil {
		t.Error("DecodeAny(missing) expected error, got nil")
	}
}
