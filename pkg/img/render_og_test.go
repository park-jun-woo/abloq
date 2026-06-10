//ff:func feature=image type=generator control=sequence
//ff:what RenderOG가 1200×630 캔버스에 흰 배경(모서리)과 제목 픽셀(중앙부 비백색)을 그리는지 검증
package img

import "testing"

func TestRenderOG(t *testing.T) {
	m, err := RenderOG(OGSpec{Title: "Hello abloq", Brand: "Test"})
	if err != nil {
		t.Fatalf("RenderOG: %v", err)
	}
	if b := m.Bounds(); b.Dx() != 1200 || b.Dy() != 630 {
		t.Fatalf("bounds = %v, want 1200x630", b)
	}
	if r, g, bl, _ := m.At(2, 2).RGBA(); r != 0xffff || g != 0xffff || bl != 0xffff {
		t.Errorf("corner pixel = %v, want white", m.At(2, 2))
	}
	if !hasInk(m) {
		t.Error("no non-white pixel found — title was not drawn")
	}
	noBrand, err := RenderOG(OGSpec{Title: "Solo"})
	if err != nil || !hasInk(noBrand) {
		t.Errorf("brand-less render failed: err %v", err)
	}
	if _, err := RenderOG(OGSpec{Title: "x", FontPath: "/nonexistent.ttf"}); err == nil {
		t.Error("RenderOG with missing font expected error, got nil")
	}
}
