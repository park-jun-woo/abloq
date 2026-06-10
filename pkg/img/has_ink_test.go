//ff:func feature=image type=parser control=iteration dimension=1
//ff:what 테스트 헬퍼 — 이미지 전체를 단일 인덱스 루프로 훑어 비백색 픽셀 존재 여부 탐색
package img

import "image"

func hasInk(m image.Image) bool {
	w, h := m.Bounds().Dx(), m.Bounds().Dy()
	for i := 0; i < w*h; i++ {
		r, g, b, _ := m.At(m.Bounds().Min.X+i%w, m.Bounds().Min.Y+i/w).RGBA()
		if r != 0xffff || g != 0xffff || b != 0xffff {
			return true
		}
	}
	return false
}
