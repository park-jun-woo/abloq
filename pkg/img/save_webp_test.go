//ff:func feature=image type=generator control=sequence
//ff:what SaveWebP가 상위 디렉토리를 만들며 WebP를 기록하고 x/image/webp로 재디코드되는지 검증
package img

import (
	"image"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/image/webp"
)

func TestSaveWebP(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sub", "out.webp")
	if err := SaveWebP(path, image.NewRGBA(image.Rect(0, 0, 6, 3))); err != nil {
		t.Fatalf("SaveWebP: %v", err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer f.Close()
	m, err := webp.Decode(f)
	if err != nil {
		t.Fatalf("webp decode: %v", err)
	}
	if b := m.Bounds(); b.Dx() != 6 || b.Dy() != 3 {
		t.Errorf("bounds = %v, want 6x3", b)
	}
	if err := SaveWebP(filepath.Join(path, "under-file.webp"), image.NewRGBA(image.Rect(0, 0, 1, 1))); err == nil {
		t.Error("SaveWebP under a regular file expected error, got nil")
	}
	if err := SaveWebP(filepath.Join(dir, "isdir.webp")+string(os.PathSeparator), image.NewRGBA(image.Rect(0, 0, 1, 1))); err == nil {
		t.Error("SaveWebP with trailing separator expected create error, got nil")
	}
}
