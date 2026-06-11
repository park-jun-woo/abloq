//ff:func feature=image type=client control=sequence
//ff:what stub 생성 1발 — Err가 있으면 그대로 실패, 아니면 고정 픽스처 단색 이미지 반환 (실 네트워크 0)
package ogprovider

import (
	"context"
	"image"
	"image/draw"
)

// Generate returns the fixed fixture canvas (or the configured error).
func (s Stub) Generate(ctx context.Context, prompt string) (image.Image, error) {
	if s.Err != nil {
		return nil, s.Err
	}
	w, h := s.W, s.H
	if w <= 0 || h <= 0 {
		w, h = 1024, 1024
	}
	m := image.NewNRGBA(image.Rect(0, 0, w, h))
	draw.Draw(m, m.Bounds(), image.NewUniform(s.Color), image.Point{}, draw.Src)
	return m, nil
}
