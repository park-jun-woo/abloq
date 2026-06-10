//ff:func feature=image type=generator control=sequence
//ff:what 이미지를 WebP(VP8L 무손실, 순수 Go 인코더)로 기록 — 상위 디렉토리 생성 포함
package img

import (
	"image"
	"os"
	"path/filepath"

	"github.com/HugoSmits86/nativewebp"
)

// SaveWebP encodes m as lossless WebP at path.
func SaveWebP(path string, m image.Image) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := nativewebp.Encode(f, m, nil); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
