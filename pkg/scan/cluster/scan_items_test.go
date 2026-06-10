//ff:func feature=scan type=generator control=sequence topic=cluster
//ff:what scanItems가 위반 글만 글당 1건으로 조립하고 무위반 글을 건너뛰는지 검증
package cluster

import "testing"

func TestScanItems(t *testing.T) {
	b := testBlog()
	posts := []post{
		{Section: "tech", Slug: "ok", Date: "2026-01-05", Tags: []string{"geo", "abloq"}, Outlinks: []string{"tech/thin", "tech/x"}},
		{Section: "tech", Slug: "thin", Date: "2026-01-04", Tags: []string{"geo", "abloq"}, Outlinks: []string{"tech/ok"}},
	}
	tags := map[string]int64{"geo": 2, "abloq": 2}
	inlinks := map[string]int64{"tech/ok": 1, "tech/thin": 1, "tech/x": 1}
	items := scanItems(posts, b, "ko", tags, inlinks, edgeSet(posts))
	if len(items) != 1 || items[0].Slug != "thin" || items[0].Kind != "cluster" {
		t.Fatalf("items = %+v, want only thin", items)
	}
}
