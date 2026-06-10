//ff:func feature=gate type=parser control=sequence topic=evidence
//ff:what newCitations 케이스 — Base 없으면 전부 신규, Base==Doc은 0건, Base에 있던 URL은 제외
package gate

import "testing"

func TestNewCitations(t *testing.T) {
	b := loadGateBlog(t)
	a := artFromContent(t, b, "Old [a](https://x.test/a) and new [b](https://x.test/b).\n")
	if got := newCitations(a); len(got) != 2 {
		t.Errorf("nil baseline: want every citation as new, got %+v", got)
	}
	a.Base = a.Doc
	if got := newCitations(a); len(got) != 0 {
		t.Errorf("unchanged article: want 0 new citations, got %+v", got)
	}
	a.Base = ParseArticle(b, "en", "Old [a](https://x.test/a) only.\n")
	got := newCitations(a)
	if len(got) != 1 || got[0].URL != "https://x.test/b" {
		t.Errorf("want only the added URL b, got %+v", got)
	}
}
