//ff:func feature=cli type=command control=sequence
//ff:what 테스트 픽스처 — 임시 디렉토리에 4×4 단색 PNG를 만들어 image convert 테스트에 제공
package main

import (
	"image"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func writePNGFixture(t *testing.T, dir string) string {
	t.Helper()
	src := filepath.Join(dir, "in.png")
	f, err := os.Create(src)
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer f.Close()
	if err := png.Encode(f, image.NewRGBA(image.Rect(0, 0, 4, 4))); err != nil {
		t.Fatalf("png encode: %v", err)
	}
	return src
}
