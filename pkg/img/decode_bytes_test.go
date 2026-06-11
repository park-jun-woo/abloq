//ff:func feature=image type=parser control=sequence
//ff:what DecodeBytes가 PNG 바이트를 디코드하고 비이미지 바이트에 에러를 내는지 검증
package img

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"testing"
)

func TestDecodeBytes(t *testing.T) {
	m := image.NewNRGBA(image.Rect(0, 0, 8, 4))
	draw.Draw(m, m.Bounds(), image.NewUniform(color.NRGBA{0, 128, 255, 255}), image.Point{}, draw.Src)
	var buf bytes.Buffer
	if err := png.Encode(&buf, m); err != nil {
		t.Fatalf("png encode: %v", err)
	}
	got, err := DecodeBytes(buf.Bytes())
	if err != nil {
		t.Fatalf("DecodeBytes: %v", err)
	}
	if b := got.Bounds(); b.Dx() != 8 || b.Dy() != 4 {
		t.Errorf("bounds = %v, want 8x4", b)
	}
	if _, err := DecodeBytes([]byte("not an image")); err == nil {
		t.Error("DecodeBytes(garbage) expected error, got nil")
	}
}
