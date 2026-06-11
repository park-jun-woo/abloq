//ff:func feature=image type=generator control=sequence
//ff:what 임의 크기 이미지를 w×h로 센터 크롭+리사이즈(CatmullRom) — 비율 보존 cover 크롭, AI 배경→1200×630 규격화
//ff:why ResizeMax는 축소 전용(비율 그대로)이라 목표 비율 강제가 불가 — provider 출력 규격화엔 cover 크롭이 따로 필요 (BUG002)
package img

import (
	"image"

	xdraw "golang.org/x/image/draw"
)

// CropCenter scales src to cover w×h and center-crops the excess: the largest
// centered source rectangle with the target aspect ratio is resampled to
// exactly w×h.
func CropCenter(src image.Image, w, h int) *image.RGBA {
	b := src.Bounds()
	cw, ch := b.Dx(), b.Dx()*h/w
	if ch > b.Dy() {
		cw, ch = b.Dy()*w/h, b.Dy()
	}
	x0 := b.Min.X + (b.Dx()-cw)/2
	y0 := b.Min.Y + (b.Dy()-ch)/2
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	xdraw.CatmullRom.Scale(dst, dst.Bounds(), src, image.Rect(x0, y0, x0+cw, y0+ch), xdraw.Src, nil)
	return dst
}
