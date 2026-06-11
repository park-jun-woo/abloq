//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what mode 1개에 대해 Build 산출물이 3종이고 static/llms.txt가 없는지 검증하는 테스트 헬퍼
package gen

import (
	"path/filepath"
	"testing"
)

func checkBuildExcludesLlms(t *testing.T, mode string) {
	t.Helper()
	b := loadGoldenBlog(t)
	b.Geo.LlmsTxt.Mode = mode
	outs := Build(filepath.Join("testdata", "golden"), b)
	if len(outs) != 3 {
		t.Fatalf("mode %s: want 3 outputs without llms.txt, got %d", mode, len(outs))
	}
	for _, o := range outs {
		if o.Path == "static/llms.txt" {
			t.Errorf("mode %s: llms.txt must be excluded from Build outputs", mode)
		}
	}
}
