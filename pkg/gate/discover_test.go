//ff:func feature=gate type=frame control=iteration dimension=1
//ff:what Discover가 미니 저장소의 전 언어·섹션 글을 언어 순서대로 수집하는지 검증
package gate

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestDiscover(t *testing.T) {
	dir := filepath.Join("testdata", "repo-pass")
	b, diags, err := blogyaml.Load(filepath.Join(dir, "blog.yaml"))
	if err != nil || len(diags) != 0 {
		t.Fatalf("fixture blog.yaml: %v %v", err, diags)
	}
	arts := Discover(dir, b)
	if len(arts) != 2 {
		t.Fatalf("want 2 articles, got %d", len(arts))
	}
	for i, want := range []string{"ko", "en"} {
		if arts[i].Lang != want || arts[i].Slug != "hello" {
			t.Errorf("arts[%d] = %s/%s, want %s/hello", i, arts[i].Lang, arts[i].Slug, want)
		}
	}
}
