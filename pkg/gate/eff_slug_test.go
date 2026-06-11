//ff:func feature=gate type=rule control=sequence
//ff:what effSlug가 front matter slug 우선, 없으면 파일 어간을 반환하는지 검증
package gate

import "testing"

func TestEffSlug(t *testing.T) {
	withFM := &Article{Slug: "file-stem", Doc: &Doc{FrontMatter: "slug: \"override\""}}
	if got := EffSlug(withFM); got != "override" {
		t.Errorf("EffSlug = %q, want override", got)
	}
	without := &Article{Slug: "file-stem", Doc: &Doc{FrontMatter: "title: x"}}
	if got := EffSlug(without); got != "file-stem" {
		t.Errorf("EffSlug = %q, want file-stem", got)
	}
}
