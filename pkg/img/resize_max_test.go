//ff:func feature=image type=generator control=sequence
//ff:what ResizeMax가 maxW 초과 이미지를 비율 유지 축소하고 이하 이미지는 원본을 반환하는지 검증
package img

import (
	"image"
	"testing"
)

func TestResizeMax(t *testing.T) {
	wide := image.NewRGBA(image.Rect(0, 0, 200, 100))
	got := ResizeMax(wide, 100)
	if b := got.Bounds(); b.Dx() != 100 || b.Dy() != 50 {
		t.Errorf("resized bounds = %v, want 100x50", b)
	}
	small := image.NewRGBA(image.Rect(0, 0, 50, 50))
	if ResizeMax(small, 100) != image.Image(small) {
		t.Error("image under maxW must be returned unchanged")
	}
}
