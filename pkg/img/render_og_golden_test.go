//ff:func feature=image type=generator control=sequence
//ff:what RenderOG+SaveWebP 출력이 오버레이 추출 전 박제한 골든 바이트와 동일한지 검증 — 텍스트 합성 추출의 회귀 0 보증
package img

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRenderOGGolden(t *testing.T) {
	want, err := os.ReadFile(filepath.Join("testdata", "og_golden.webp"))
	if err != nil {
		t.Fatalf("read golden: %v", err)
	}
	out := filepath.Join(t.TempDir(), "og.webp")
	spec := OGSpec{Title: "Agentic blog Quest — the blog that agents run", Brand: "parkjunwoo.com", Out: out}
	if err := OG(spec); err != nil {
		t.Fatalf("OG: %v", err)
	}
	got, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("read output: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("RenderOG output differs from the pre-extraction golden (%d vs %d bytes)", len(got), len(want))
	}
}
