//ff:func feature=image type=parser control=sequence
//ff:what 이미지 바이트를 디코드 — png/jpeg/gif/webp (포맷 등록은 decode_any.go의 blank import 재사용)
//ff:why DecodeAny는 파일 경로 시그니처(os.Open)라 API 응답 바이트에 못 쓴다 — provider 응답용 바이트 입구를 분리 (BUG002)
package img

import (
	"bytes"
	"image"
)

// DecodeBytes decodes an in-memory image with the formats registered by this
// package (see decode_any.go).
func DecodeBytes(data []byte) (image.Image, error) {
	m, _, err := image.Decode(bytes.NewReader(data))
	return m, err
}
