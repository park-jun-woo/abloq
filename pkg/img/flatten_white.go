//ff:func feature=image type=generator control=sequence
//ff:what 이미지를 흰색 배경 RGBA로 평탄화 — 투명 픽셀(RGBA 원본)을 흰색 위에 합성
package img

import (
	"image"
	"image/color"
	"image/draw"
)

// FlattenWhite composites src over a white background, removing alpha.
func FlattenWhite(src image.Image) *image.RGBA {
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(dst, dst.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)
	draw.Draw(dst, dst.Bounds(), src, b.Min, draw.Over)
	return dst
}
