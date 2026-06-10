//ff:func feature=image type=generator control=sequence
//ff:what FlattenWhite가 투명 픽셀을 흰색으로, 불투명 픽셀은 원색 그대로 합성하는지 검증
package img

import (
	"image"
	"image/color"
	"testing"
)

func TestFlattenWhite(t *testing.T) {
	src := image.NewNRGBA(image.Rect(0, 0, 2, 1))
	src.SetNRGBA(0, 0, color.NRGBA{0, 0, 0, 0})       // transparent
	src.SetNRGBA(1, 0, color.NRGBA{255, 0, 0, 255}) // opaque red
	got := FlattenWhite(src)
	r, g, b, _ := got.At(0, 0).RGBA()
	if r != 0xffff || g != 0xffff || b != 0xffff {
		t.Errorf("transparent pixel = %v, want white", got.At(0, 0))
	}
	r, g, b, _ = got.At(1, 0).RGBA()
	if r != 0xffff || g != 0 || b != 0 {
		t.Errorf("opaque pixel = %v, want red", got.At(1, 0))
	}
}
