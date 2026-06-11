//ff:func feature=image type=generator control=sequence
//ff:what OG 이미지(1200×630)를 렌더 — 흰 캔버스를 만들어 OverlayText로 제목/브랜드를 합성 (현행 바이트 동일)
package img

import (
	"image"
	"image/color"
	"image/draw"
)

// RenderOG draws spec onto a 1200x630 white canvas via OverlayText. The title
// wraps to fit, the optional brand line sits below in the accent color.
func RenderOG(spec OGSpec) (image.Image, error) {
	const w, h = 1200, 630
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(dst, dst.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)
	if err := OverlayText(dst, spec); err != nil {
		return nil, err
	}
	return dst, nil
}
