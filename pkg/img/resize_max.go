//ff:func feature=image type=generator control=sequence
//ff:what 가로 maxW 초과 이미지를 비율 유지 축소(CatmullRom) — 이하면 원본 그대로 반환
package img

import (
	"image"

	xdraw "golang.org/x/image/draw"
)

// ResizeMax scales src down to maxW pixels wide when it is wider.
func ResizeMax(src image.Image, maxW int) image.Image {
	b := src.Bounds()
	if maxW <= 0 || b.Dx() <= maxW {
		return src
	}
	h := b.Dy() * maxW / b.Dx()
	dst := image.NewRGBA(image.Rect(0, 0, maxW, h))
	xdraw.CatmullRom.Scale(dst, dst.Bounds(), src, b, xdraw.Src, nil)
	return dst
}
