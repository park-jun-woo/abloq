//ff:func feature=image type=parser control=sequence
//ff:what 이미지 파일 1개를 디코드 — png/jpeg/gif/webp 지원 (외부 바이너리 없이 순수 Go)
package img

import (
	"image"
	_ "image/gif"  // register gif
	_ "image/jpeg" // register jpeg
	_ "image/png"  // register png
	"os"

	_ "golang.org/x/image/webp" // register webp (decode)
)

// DecodeAny opens and decodes path with the registered formats.
func DecodeAny(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	m, _, err := image.Decode(f)
	return m, err
}
