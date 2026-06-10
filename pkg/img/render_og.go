//ff:func feature=image type=generator control=iteration dimension=1
//ff:what OG 이미지(1200×630)를 렌더 — 흰 배경, 제목 줄바꿈 중앙 정렬, 하단 액센트 브랜드 라인
package img

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// RenderOG draws spec onto a 1200x630 canvas. The title wraps to fit, the
// optional brand line sits below in the accent color.
func RenderOG(spec OGSpec) (image.Image, error) {
	const w, h, lineH, brandGap = 1200, 630, 84, 60
	titleFace, err := LoadFace(spec.FontPath, 64)
	if err != nil {
		return nil, err
	}
	brandFace, err := LoadFace(spec.FontPath, 34)
	if err != nil {
		return nil, err
	}
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(dst, dst.Bounds(), image.NewUniform(color.White), image.Point{}, draw.Src)
	lines := WrapText(titleFace, spec.Title, w-160)
	total := len(lines) * lineH
	if spec.Brand != "" {
		total += brandGap + 34
	}
	y := (h-total)/2 + 56
	d := font.Drawer{Dst: dst, Src: image.NewUniform(color.RGBA{30, 30, 30, 255}), Face: titleFace}
	for _, line := range lines {
		d.Dot = fixed.P((w-font.MeasureString(titleFace, line).Ceil())/2, y)
		d.DrawString(line)
		y += lineH
	}
	if spec.Brand != "" {
		b := font.Drawer{Dst: dst, Src: image.NewUniform(color.RGBA{38, 127, 192, 255}), Face: brandFace}
		b.Dot = fixed.P((w-font.MeasureString(brandFace, spec.Brand).Ceil())/2, y+brandGap-lineH/2)
		b.DrawString(spec.Brand)
	}
	return dst, nil
}
