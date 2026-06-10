//ff:func feature=image type=generator control=sequence
//ff:what Convert가 PNG를 WebP로 변환(흰 배경 평탄화 + 축소)하고 크기 보고를 채우는지 검증
package img

import (
	"image/color"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/image/webp"
)

func TestConvert(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "in.png")
	writeTestPNG(t, src, 200, 100, color.NRGBA{0, 128, 255, 255})
	dst := filepath.Join(dir, "out", "in.webp")
	res, err := Convert(src, dst, 100)
	if err != nil {
		t.Fatalf("Convert: %v", err)
	}
	if res.SrcBytes <= 0 || res.DstBytes <= 0 || res.Dst != dst {
		t.Errorf("ConvertResult = %+v, want positive sizes", res)
	}
	f, err := os.Open(dst)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer f.Close()
	m, err := webp.Decode(f)
	if err != nil {
		t.Fatalf("webp decode: %v", err)
	}
	if b := m.Bounds(); b.Dx() != 100 || b.Dy() != 50 {
		t.Errorf("converted bounds = %v, want 100x50", b)
	}
	if _, err := Convert(filepath.Join(dir, "missing.png"), dst, 100); err == nil {
		t.Error("Convert(missing src) expected error, got nil")
	}
	if _, err := Convert(src, filepath.Join(src, "blocked.webp"), 100); err == nil {
		t.Error("Convert with dst under a regular file expected error, got nil")
	}
}
