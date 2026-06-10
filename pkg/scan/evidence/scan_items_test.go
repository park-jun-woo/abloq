//ff:func feature=scan type=generator control=sequence topic=evidence
//ff:what scanItems가 주장 보유·rot 보유 글만 항목으로 만들고 깨끗한 글을 제외하는지 검증
package evidence

import (
	"testing"

	"github.com/park-jun-woo/abloq/pkg/gate"
)

func TestScanItems(t *testing.T) {
	b := testBlog(t)
	claimsArt := testArticle(t, b, "---\ntitle: C\n---\n\n처리량이 40% 증가했다.\n")
	rotArt := testArticle(t, b, "---\ntitle: R\n---\n\n[ref](https://gone.example/x)\n")
	rotArt.Slug = "rot-post"
	cleanArt := testArticle(t, b, "---\ntitle: K\n---\n\n깨끗한 본문.\n")
	cleanArt.Slug = "clean-post"
	checks := []Check{{URL: "https://gone.example/x", Lang: "ko", Section: "tech", Slug: "rot-post",
		Status: "hard", ConsecutiveFailures: 3}}
	items := scanItems([]*gate.Article{claimsArt, rotArt, cleanArt}, checks)
	if len(items) != 2 {
		t.Fatalf("want 2 items, got %d: %+v", len(items), items)
	}
	if items[0].Slug != "fixture" || items[0].Payload["claims"] == "" {
		t.Errorf("claims item: %+v", items[0])
	}
	if items[1].Slug != "rot-post" || items[1].Payload["rot_urls"] == "" {
		t.Errorf("rot item: %+v", items[1])
	}
}
