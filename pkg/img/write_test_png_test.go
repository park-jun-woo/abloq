//ff:func feature=image type=generator control=sequence
//ff:what 테스트 픽스처 — 지정 크기·단색 PNG 파일을 기록하는 헬퍼
package img

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"testing"
)

func writeTestPNG(t *testing.T, path string, w, h int, c color.Color) {
	t.Helper()
	m := image.NewNRGBA(image.Rect(0, 0, w, h))
	draw.Draw(m, m.Bounds(), image.NewUniform(c), image.Point{}, draw.Src)
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("create %s: %v", path, err)
	}
	defer f.Close()
	if err := png.Encode(f, m); err != nil {
		t.Fatalf("png encode: %v", err)
	}
}
