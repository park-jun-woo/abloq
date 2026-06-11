//ff:func feature=gen type=parser control=iteration dimension=1
//ff:what Collect가 골든 콘텐츠 트리에서 all 스코프 언어·선언 섹션의 발행 글만(draft 제외) 수집하는지 검증
package llms

import (
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/blogyaml"
)

func TestCollect(t *testing.T) {
	root := filepath.Join("..", "testdata", "golden")
	b := &blogyaml.Blog{
		Languages: []string{"ko", "en"},
		Sections:  []string{"opinion", "tech"},
		Geo:       blogyaml.Geo{LlmsTxt: blogyaml.LlmsTxtSpec{Languages: []string{"all"}}},
	}
	posts := Collect(root, b)
	if len(posts) != 4 {
		t.Fatalf("want 4 published posts, got %d: %+v", len(posts), posts)
	}
	for _, p := range posts {
		if p.Slug == "draft-post" {
			t.Errorf("draft post must be excluded: %+v", p)
		}
	}
}
