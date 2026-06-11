//ff:func feature=image type=generator control=iteration dimension=1
//ff:what CropCenter가 가로/세로 초과·축소 입력 전부를 정확히 w×h로 만들고 중앙부를 보존하는지 검증
package img

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
)

func TestCropCenter(t *testing.T) {
	// wide source: 2000x630 — sides cropped; tall source: 1200x2000 — top/bottom
	// cropped; small source: 300x100 — upscaled to cover.
	for _, size := range []image.Point{{2000, 630}, {1200, 2000}, {300, 100}} {
		src := image.NewNRGBA(image.Rect(0, 0, size.X, size.Y))
		draw.Draw(src, src.Bounds(), image.NewUniform(color.NRGBA{200, 10, 10, 255}), image.Point{}, draw.Src)
		got := CropCenter(src, 1200, 630)
		if b := got.Bounds(); b.Dx() != 1200 || b.Dy() != 630 {
			t.Errorf("CropCenter(%v) bounds = %v, want 1200x630", size, b)
		}
		if r, _, _, _ := got.At(600, 315).RGBA(); r < 0xb000 {
			t.Errorf("CropCenter(%v) center pixel lost source color: %v", size, got.At(600, 315))
		}
	}
	// center preservation: a wide source with a distinct center column keeps it.
	src := image.NewNRGBA(image.Rect(0, 0, 2400, 630))
	draw.Draw(src, src.Bounds(), image.NewUniform(color.NRGBA{255, 255, 255, 255}), image.Point{}, draw.Src)
	draw.Draw(src, image.Rect(1180, 0, 1220, 630), image.NewUniform(color.NRGBA{0, 0, 0, 255}), image.Point{}, draw.Src)
	got := CropCenter(src, 1200, 630)
	if r, g, b, _ := got.At(600, 315).RGBA(); r > 0x4000 || g > 0x4000 || b > 0x4000 {
		t.Errorf("center stripe not preserved: %v", got.At(600, 315))
	}
}
