//ff:func feature=scan type=parser control=sequence topic=cluster
//ff:what postFromArticle이 태그·날짜·slug 오버라이드를 디코드하고 draft를 제외하는지 검증
package cluster

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/abloq/pkg/gate"
)

func TestPostFromArticle(t *testing.T) {
	b := testBlog()
	a := &gate.Article{Lang: "ko", Section: "tech", Slug: "stem", Doc: &gate.Doc{
		FrontMatter: "title: T\ndate: 2026-01-05\nslug: override\ntags: [geo, abloq]\n",
		Body:        "[t](/tech/thin/)\n",
	}}
	p, ok := postFromArticle(b, "ko", a)
	if !ok {
		t.Fatal("published article must resolve")
	}
	if p.Slug != "override" || p.Date != "2026-01-05" || p.Section != "tech" {
		t.Errorf("post = %+v", p)
	}
	if !reflect.DeepEqual(p.Tags, []string{"geo", "abloq"}) || !reflect.DeepEqual(p.Outlinks, []string{"tech/thin"}) {
		t.Errorf("tags/outlinks = %v / %v", p.Tags, p.Outlinks)
	}
	draft := &gate.Article{Lang: "ko", Section: "tech", Slug: "d", Doc: &gate.Doc{FrontMatter: "draft: true\n"}}
	if _, ok := postFromArticle(b, "ko", draft); ok {
		t.Error("draft must be excluded")
	}
	bare := &gate.Article{Lang: "ko", Section: "tech", Slug: "bare", Doc: &gate.Doc{FrontMatter: "title: B\n"}}
	p, ok = postFromArticle(b, "ko", bare)
	if !ok || p.Slug != "bare" || len(p.Tags) != 0 {
		t.Errorf("bare = %+v, ok=%v", p, ok)
	}
}
