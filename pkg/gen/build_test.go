//ff:func feature=gen type=generator control=iteration dimension=1
//ff:what 골든 blog.yaml + 콘텐츠 트리에서 Build 산출물 4종이 기대 스냅샷(want/)과 바이트 일치하는지 검증
package gen

import (
	"path/filepath"
	"testing"
)

func TestBuild(t *testing.T) {
	dir := filepath.Join("testdata", "golden")
	outs := Build(dir, loadGoldenBlog(t))
	wantPaths := []string{"hugo.toml", "static/robots.txt", "static/llms.txt", "data/jsonld.json"}
	if len(outs) != len(wantPaths) {
		t.Fatalf("want %d outputs, got %d", len(wantPaths), len(outs))
	}
	for i, o := range outs {
		if o.Path != wantPaths[i] {
			t.Errorf("outs[%d].Path = %q, want %q", i, o.Path, wantPaths[i])
		}
		checkGoldenOutput(t, dir, o)
	}
}
