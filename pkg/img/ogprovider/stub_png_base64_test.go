//ff:func feature=image type=client control=sequence
//ff:what 테스트 픽스처 — w×h 단색 PNG를 base64 문자열로 만들어 스텁 서버 응답에 제공
package ogprovider

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"testing"
)

func stubPNGBase64(t *testing.T, w, h int) string {
	t.Helper()
	m := image.NewNRGBA(image.Rect(0, 0, w, h))
	draw.Draw(m, m.Bounds(), image.NewUniform(color.NRGBA{40, 90, 200, 255}), image.Point{}, draw.Src)
	var buf bytes.Buffer
	if err := png.Encode(&buf, m); err != nil {
		t.Fatalf("png encode: %v", err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
