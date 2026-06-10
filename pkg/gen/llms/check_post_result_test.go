//ff:func feature=gen type=parser control=sequence
//ff:what postFromEntry 반환값(ok/slug/title)을 기대값과 비교 검증
package llms

import "testing"

func checkPostResult(t *testing.T, p Post, ok, wantOK bool, wantSlug, wantTitle string) {
	t.Helper()
	if ok != wantOK {
		t.Fatalf("postFromEntry ok = %v, want %v (post %+v)", ok, wantOK, p)
	}
	if p.Slug != wantSlug {
		t.Errorf("slug = %q, want %q", p.Slug, wantSlug)
	}
	if p.Title != wantTitle {
		t.Errorf("title = %q, want %q", p.Title, wantTitle)
	}
}
