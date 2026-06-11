//ff:func feature=image type=generator control=iteration dimension=1
//ff:what 임의 배경 캔버스에 제목/브랜드 텍스트를 합성 — 제목 줄바꿈 중앙 정렬, 하단 액센트 브랜드 라인 (결정론)
//ff:why RenderOG의 흰 캔버스 직결 합성을 추출 — AI 배경(--overlay) 위에도 같은 결정론 텍스트를 올린다 (BUG002)
package img

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// OverlayText draws spec's title (wrapped, centered) and the optional brand
// line onto dst. dst can be any RGBA canvas — the white OG card and AI
// backgrounds share this one deterministic composition.
func OverlayText(dst draw.Image, spec OGSpec) error {
	const lineH, brandGap = 84, 60
	w, h := dst.Bounds().Dx(), dst.Bounds().Dy()
	titleFace, err := LoadFace(spec.FontPath, 64)
	if err != nil {
		return err
	}
	brandFace, err := LoadFace(spec.FontPath, 34)
	if err != nil {
		return err
	}
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
	return nil
}
