//ff:func feature=image type=client control=sequence
//ff:what 테스트 stub Generate — err 지정 시 실패, 아니면 w×h 단색 이미지 반환
package img

import (
	"context"
	"image"
	"image/draw"
)

func (s stubOGProvider) Generate(ctx context.Context, prompt string) (image.Image, error) {
	if s.err != nil {
		return nil, s.err
	}
	m := image.NewNRGBA(image.Rect(0, 0, s.w, s.h))
	draw.Draw(m, m.Bounds(), image.NewUniform(s.c), image.Point{}, draw.Src)
	return m, nil
}
