//ff:func feature=scan type=parser control=sequence topic=evidence
//ff:what collectCitations이 전 글의 인용을 글 순서대로 모으고, 같은 URL이라도 글이 다르면 별개 cite로 두는지 검증
package evidence

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/gate"
)

func TestCollectCitations(t *testing.T) {
	b := testBlog(t)
	a1 := testArticle(t, b, "---\ntitle: A\n---\n\n[ref](https://shared.example/x)\n")
	a2 := testArticle(t, b, "---\ntitle: B\n---\n\n[ref](https://shared.example/x)\n")
	a2.Slug = "other"
	cites := collectCitations([]*gate.Article{a1, a2})
	if len(cites) != 2 {
		t.Fatalf("want 2 cites (same URL, two articles), got %d", len(cites))
	}
	if cites[0].Slug == cites[1].Slug {
		t.Errorf("cites must keep per-article keys: %+v", cites)
	}
}
