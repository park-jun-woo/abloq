//ff:func feature=gate type=rule control=sequence
//ff:what slugMismatchDiag가 front matter slug 불일치를 검출하고 일치 시 nil을 반환하는지 검증
package gate

import "testing"

func TestSlugMismatchDiag(t *testing.T) {
	group := []*Article{
		{Lang: "ko", Slug: "a", Path: "ko/a.md", Doc: &Doc{FrontMatter: "slug: \"a\""}},
		{Lang: "en", Slug: "a", Path: "en/a.md", Doc: &Doc{FrontMatter: "slug: \"b\""}},
	}
	d := slugMismatchDiag(group)
	if d == nil || d.Rule != "slug-consistency" || d.File != "en/a.md" {
		t.Fatalf("diag = %v, want slug-consistency on en/a.md", d)
	}
	same := []*Article{
		{Lang: "ko", Slug: "a", Doc: &Doc{FrontMatter: ""}},
		{Lang: "en", Slug: "a", Doc: &Doc{FrontMatter: "slug: \"a\""}},
	}
	if got := slugMismatchDiag(same); got != nil {
		t.Errorf("matching slugs: want nil, got %v", got)
	}
}
